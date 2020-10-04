package gocli

import (
	"context"
)

type Action struct {
	Name	string
	Method 	ActionMethod
	Options []*Option
}

type ActionMethod interface {
	Description(cli CLI) string
	Run(ctx context.Context, cli CLI) error
}

func (c *CLI) AddAction(name string, handler ActionMethod) *Action {
	c.Actions = append(c.Actions, &Action{
		Name: name,
		Method: handler,
	})
	return c.Actions[len(c.Actions)-1]
}

