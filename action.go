package cli

import (
	"context"
	"log"
)

type Action struct {
	Name	string
	Args 	[]string
	Method 	ActionMethod
	options []*Option
}

type ActionMethod interface {
	Description(cli CLI) string
	Run(ctx context.Context, cli CLI) error
}

func AddAction(handler ActionMethod, name string, args ...string) *Action {
	if _, err := findAction(name); err == nil {
		log.Fatal("cli: duplicate method name '" + name + "'")
	}
	c.actions = append(c.actions, &Action{
		Name: name,
		Method: handler,
		Args: args,
	})
	return c.actions[len(c.actions)-1]
}

