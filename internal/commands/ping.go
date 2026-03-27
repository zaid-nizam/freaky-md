package commands

import (
	"fmt"
	"time"

	"go.mau.fi/whatsmeow/proto/waE2E"
)

func init() {
	Register(Command{
		Name:        "ping",
		Description: "Checks bot latency",
		Aliases:     []string{"p"},
		Execute:     handlePing,
	})
}

func handlePing(ctx *CommandContext) {
	latency := time.Since(ctx.Event.Info.Timestamp)
	res := fmt.Sprintf("🏓 *Pong!*\n_Latency: %dms_", latency.Milliseconds())

	_, err := ctx.Client.SendMessage(ctx.Context, ctx.Event.Info.Chat, &waE2E.Message{
		Conversation: &res,
	})

	if err != nil {
		fmt.Printf("Ping error: %v\n", err)
	}
}

