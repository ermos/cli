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
	Options 	map[string][]string
	actions 	[]*Action
	options 	[]*Option
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
	// Global Help
	if len(c.Args) < 1 || c.Args[0] == "--help" {
		showGlobalHelp()
		return
	}
	// Before All Middleware
	err = callMiddleware(ctx, beforeAllMiddleware)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Global Options
	err = findOptions(c.options)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Before Middleware
	err = callMiddleware(ctx, beforeMiddleware)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(c.Args) == 0 {
		return
	}
	// Process
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
	// After Middleware
	err = callMiddleware(ctx, afterMiddleware)
	if err != nil {
		fmt.Println(err.Error())
		return
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

func removeArgs(from, length int) {
	start := c.Args[:from]
	if from != len(c.Args) {
		end := c.Args[from+length:]
		start = append(start, end...)
	}
	c.Args = start
}

func findOptions(o []*Option) error {
	result := make(map[string][]string)
	for key, value := range c.Options {
		result[key] = value
	}
	possibility := make(map[string]*Option)
	for _, opt := range o {
		possibility["--" + opt.Name] = opt
		if opt.ShortName != "" {
			possibility["-" + opt.ShortName] = opt
		}
	}
	length := len(c.Args)
	totalArg := -1
	for i := 0; i < length; i++ {
		if possibility[c.Args[i]] == nil {
			totalArg = i
			break
		}
		var values []string
		currI := i
		for a := 0; a < len(possibility[c.Args[currI]].ArgsType); a++ {
			if i+1 >= length {
				return errors.New(fmt.Sprintf(
					"%s: '%s' need %d arguments.\nSee '%s --help'",
					c.Name,
					c.Args[currI],
					len(possibility[c.Args[currI]].ArgsType),
					c.Name,
					))
			}
			i++
			values = append(values, c.Args[i])
		}
		result[possibility[c.Args[currI]].Name] = values
	}
	if totalArg == -1 {
		totalArg = length
	}
	if totalArg != 0 {
		removeArgs(0, totalArg)
	}
	c.Options = result
	return nil
}