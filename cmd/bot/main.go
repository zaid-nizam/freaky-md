package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"

	"freaky_md/internal/bot"
	_ "freaky_md/internal/commands"
	"freaky_md/internal/config"
)

func main() {
	cfg := config.Load("config.yaml")

	logLevel := "INFO"
	if cfg.Debug {
		logLevel = "DEBUG"
	}
	dbLog := waLog.Stdout("Database", logLevel, true)
	clientLog := waLog.Stdout("Client", logLevel, true)

	ctx := context.Background()

	dbPath := fmt.Sprintf("file:%s?_foreign_keys=on", cfg.SessionPath)
	container, err := sqlstore.New(ctx, "sqlite3", dbPath, dbLog)
	if err != nil {
		clientLog.Errorf("Critical: Failed to connect to database: %v", err)
		os.Exit(1)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		clientLog.Errorf("Critical: Failed to get device store: %v", err)
		os.Exit(1)
	}

	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(bot.Handler(client, cfg))

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())

		err = client.Connect()
		if err != nil {
			clientLog.Errorf("Critical: Initial connection failed: %v", err)
			os.Exit(1)
		}

		for evt := range qrChan {
			switch evt.Event {
			case "code":
				qr, err := qrcode.New(evt.Code, qrcode.Medium)
				if err != nil {
					clientLog.Errorf("QR Generation Error: %v", err)
					continue
				}
				fmt.Println("\n[ACTION REQUIRED] Scan the QR code below to login:")
				fmt.Println(qr.ToSmallString(false))
			case "success":
				clientLog.Infof("Login successful!")
			case "timeout":
				clientLog.Errorf("QR link timed out. Restart the bot to try again.")
				os.Exit(1)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			clientLog.Errorf("Critical: Failed to reconnect: %v", err)
			os.Exit(1)
		}
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	clientLog.Infof("Bot is online. Press Ctrl+C to stop.")
	<-stop

	clientLog.Infof("Shutting down... disconnecting from WhatsApp")
	client.Disconnect()

	time.Sleep(1 * time.Second)
	clientLog.Infof("Exited safely.")
}
