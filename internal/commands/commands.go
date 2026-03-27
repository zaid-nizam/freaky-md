package commands

import (
	"context"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

type CommandContext struct {
	Context context.Context
	Client  *whatsmeow.Client
	Event   *events.Message
	Args    []string
}

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Execute     func(ctx *CommandContext)
}

var Registry = make(map[string]Command)

func Register(cmd Command) {
	Registry[cmd.Name] = cmd
	for _, alias := range cmd.Aliases {
		Registry[alias] = cmd
	}
}
