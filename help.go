package cli

import (
	"fmt"
	"strings"
)

var tabSpace = 4

func showGlobalHelp() {
	Write(`
Usage:  %s [OPTIONS] COMMAND

%s

%s
Commands:
%s
Run '%s COMMAND --help' for more information on a command.
`,
	c.Name,
	c.Description,
	printOptions(c.options),
	printCommands(),
	c.Name,
	)
}

func printCommands() string {
	var list string
	var maxLength int
	for _, a := range c.actions {
		if len(a.Name) > maxLength {
			maxLength = len(a.Name)
		}
	}
	for _, a := range c.actions {
		list += fmt.Sprintln("  " + a.Name + strings.Repeat(" ", maxLength - len(a.Name) + tabSpace) + a.Method.Description(c))
	}
	return list
}

func showCommandHelp(a *Action) {
	var args []string
	for _, arg := range a.Args {
		args = append(args, "[" + strings.ToUpper(arg) + "]")
	}
	options := printOptions(a.options)
	argOption := " [OPTIONS]"
	if options == "" {
		argOption = ""
	}
	Write("\nUsage:  %s %s%s %s\n\n%s%s",
		c.Name,
		a.Name,
		argOption,
		strings.Join(args, " "),
		a.Method.Description(c),
		printOptions(a.options),
	)
}

func printOptions(opts []*Option) string {
	var list []string
	var maxLength, maxLengthShort int
	for _, o := range opts {
		if len(o.Name) + len(strings.Join(o.ArgsType, " ")) > maxLength {
			maxLength = len(o.Name) + len(strings.Join(o.ArgsType, " "))
		}
		if len(o.ShortName) > maxLengthShort {
			maxLengthShort = len(o.ShortName)
		}
	}
	for _, o := range opts {
		var short string
		if o.ShortName != ""  {
			short = strings.Repeat(" ", maxLengthShort - len(o.ShortName)) + "-" + o.ShortName + ", "
		} else {
			short = strings.Repeat(" ", maxLengthShort - len(o.ShortName) + 3)
		}
		list = append(list, fmt.Sprintf(
			"  %s%s %s%s",
			short,
			"--" + o.Name,
			strings.Join(o.ArgsType, " ") + strings.Repeat(" ", maxLength-len(o.Name)-len(strings.Join(o.ArgsType, " "))+tabSpace),
			o.Description,
		))
	}
	if len(list) != 0 {
		return fmt.Sprintf("\n\nOptions:\n%s", strings.Join(list, "\n"))
	}
	return ""
}