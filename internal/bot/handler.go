package bot

import (
	"context"
	"strings"

	"freaky_md/internal/commands"
	"freaky_md/internal/config"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

func Handler(client *whatsmeow.Client, cfg *config.Config) func(interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			handleMessage(client, v, cfg)
		}
	}
}

func handleMessage(client *whatsmeow.Client, evt *events.Message, cfg *config.Config) {
	if evt.Info.IsFromMe || evt.Message == nil {
		return
	}

	if cfg.AutoRead {
		client.MarkRead(
			context.Background(),
			[]types.MessageID{evt.Info.ID},
			evt.Info.Timestamp,
			evt.Info.Chat,
			evt.Info.Sender,
		)
	}

	var text string
	if conv := evt.Message.GetConversation(); conv != "" {
		text = conv
	} else if ext := evt.Message.GetExtendedTextMessage(); ext != nil {
		text = ext.GetText()
	}

	var prefixUsed string
	for _, p := range cfg.Prefixes {
		if strings.HasPrefix(strings.ToLower(text), strings.ToLower(p)) {
			prefixUsed = p
			break
		}
	}

	if prefixUsed == "" {
		return
	}

	cleanText := strings.TrimPrefix(text, prefixUsed)
	parts := strings.Fields(cleanText)
	if len(parts) == 0 {
		return
	}

	cmdName := strings.ToLower(parts[0])
	args := parts[1:]

	cmd, exists := commands.Registry[cmdName]
	if !exists {
		return
	}

	ctx := &commands.CommandContext{
		Context: context.Background(),
		Client:  client,
		Event:   evt,
		Args:    args,
		Config:  cfg,
	}

	cmd.Execute(ctx)
}

