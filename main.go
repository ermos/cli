package gocli

import (
	"context"
	"errors"
	"fmt"
	"github.com/ermos/gocli/help"
	"log"
	"os"
)

type CLI struct {
	Name 		string
	Description	string
	Actions 	[]*Action
}

func Init(name, description string) *CLI {
	return &CLI {
		Name: name,
		Description: description,
		Actions: []*Action{
			{
				Name: "help",
				Method: help.Handler{},
			},
		},
	}
}

func (c *CLI) Run() {
	var a *Action
	var err error
	ctx := context.Background()
	removeArgs(0, 1)
	if len(os.Args) < 1 {
		a, err = c.findAction("help")
		if err != nil {
			log.Fatal("cannot find help command")
		}
		if err = a.Method.Run(ctx, *c); err != nil {
			log.Fatal(err.Error())
		}
	}
	//a, err := c.findAction(os.Args[1])
}

func (c *CLI) findAction(name string) (*Action, error) {
	for _, action := range c.Actions {
		if name == action.Name {
			return action, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("'%s' is not a %s command.\nSee '%s --help'", c.Name, name, c.Name))
}

func removeArgs (from, length int) {
	start := os.Args[:from]
	if from != len(os.Args) {
		end := os.Args[from+length:]
		start = append(start, end...)
	}
	os.Args = start
}
