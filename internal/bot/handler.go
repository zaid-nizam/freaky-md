package bot

import (
	"context"
	"strings"

	"freaky_md/internal/commands"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func Handler(client *whatsmeow.Client) func(interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			handleMessage(client, v)
		}
	}
}

func handleMessage(client *whatsmeow.Client, evt *events.Message) {
	if evt.Info.IsFromMe || evt.Message == nil {
		return
	}

	var text string
	if conv := evt.Message.GetConversation(); conv != "" {
		text = conv
	} else if ext := evt.Message.GetExtendedTextMessage(); ext != nil {
		text = ext.GetText()
	}

	if !strings.HasPrefix(text, "!") {
		return
	}

	parts := strings.Fields(strings.TrimPrefix(text, "!"))
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
	}

	cmd.Execute(ctx)
}
