package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	Name 		string
	Description	string
	Current struct {
		Action 	string
	}
	args 		[]string
	Args 		[]string
	Options 	map[string][]string
	actions 	[]*Action
	options 	[]*Option
}

var c = CLI {
	Name: "cli",
	Description: "default cli description",
	args: os.Args,
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
	if len(c.args) < 1 || c.args[0] == "--help" {
		showGlobalHelp()
		return
	}
	// Get Action
	a, err = findAction(c.args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.Current.Action = a.Name
	removeArgs(0, 1)
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
	if len(c.args) == 0 {
		return
	}
	// Process
	if len(c.args) < 1 || c.args[0] != "--help" {
		// Process Options
		err = findOptions(a.options)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Process Args
		err = chekArgs(a)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.Args = c.args
		// Run Process
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
	start := c.args[:from]
	if from != len(c.args) {
		end := c.args[from+length:]
		start = append(start, end...)
	}
	c.args = start
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
	length := len(c.args)
	totalArg := -1
	for i := 0; i < length; i++ {
		if possibility[c.args[i]] == nil {
			totalArg = i
			break
		}
		var values []string
		currI := i
		for a := 0; a < len(possibility[c.args[currI]].ArgsType); a++ {
			if i+1 >= length {
				return errors.New(fmt.Sprintf(
					"%s: '%s' need %d arguments.\nSee '%s --help'",
					c.Name,
					c.args[currI],
					len(possibility[c.args[currI]].ArgsType),
					c.Name,
					))
			}
			i++
			values = append(values, c.args[i])
		}
		result[possibility[c.args[currI]].Name] = values
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

func chekArgs(a *Action) error {
	length := len(a.Args)
	isVariadic := false
	if length != 0 && strings.HasPrefix(a.Args[length-1], "...") {
		isVariadic = true
	}
	if length != len(c.args) {
		if (len(c.args) == length-1 || len(c.args) > length) && isVariadic {
			return nil
		}
		if !isVariadic {
			return errors.New(fmt.Sprintf(
				"%s: '%s' require %d arguments.\nSee '%s %s --help'",
				c.Name,
				a.Name,
				length,
				c.Name,
				a.Name,
			))
		} else {
			return errors.New(fmt.Sprintf(
				"%s: '%s' require between %d and infinite arguments.\nSee '%s %s --help'",
				c.Name,
				a.Name,
				length-1,
				c.Name,
				a.Name,
			))
		}
	}
	return nil
}