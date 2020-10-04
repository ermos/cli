package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
)

type CLI struct {
	Name 		string
	Description	string
	Args 		[]string
	actions 	[]*Action
}

var c = CLI {
	Name: "cli",
	Description: "default cli description",
	Args: os.Args,
}

func Init(name, description string) {
	c.Name = name
	c.Description = description
}

func Run() {
	var a *Action
	var err error
	ctx := context.Background()
	removeArgs(0, 1)
	if len(c.Args) < 1 {
		showGlobalHelp()
		return
	}
	a, err = findAction(c.Args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	removeArgs(0, 1)
	if len(c.Args) < 1 || c.Args[0] != "--help" {
		if err = a.Method.Run(ctx, c); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		showCommandHelp(a)
	}
}

func findAction(name string) (*Action, error) {
	for _, action := range c.actions {
		if name == action.Name {
			return action, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("%s: '%s' is not a %s command.\nSee '%s --help'", c.Name, name, c.Name, c.Name))
}

func removeArgs (from, length int) {
	start := c.Args[:from]
	if from != len(c.Args) {
		end := c.Args[from+length:]
		start = append(start, end...)
	}
	c.Args = start
}
