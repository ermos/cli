package cli

import (
	"fmt"
	"log"
)

// Argument's Type
const (
	AtString = 0
	AtInt = 1
	AtBool = 2
	AtFloat = 3
	AtDate = 4
)

var argTypeName = []string{
	"string",
	"int",
	"bool",
	"float",
	"date",
}

type Option struct {
	Name 		string
	ShortName 	string
	ArgsType	[]string
	Description string
}

func AddOption(name, description string, argsType ...int) *Option {
	o := addOption(name, description, argsType...)
	c.Options = append(c.Options, o)
	return o
}

func (a *Action) AddOption(name, description string, argsType ...int) *Option {
	o := addOption(name, description, argsType...)
	a.Options = append(a.Options, o)
	return o
}

func addOption (name, description string, argsType ...int) *Option {
	o := new(Option)
	o.Name = name
	o.Description = description
	length := len(argTypeName)-1
	var args []string
	for _, argID := range argsType {
		if argID > length {
			log.Fatal(fmt.Sprintf("cli: unknow argument's type ID %d", argID))
		}
		args = append(args, argTypeName[argID])
	}
	o.ArgsType = args
	return o
}

func (o *Option) AddShortName(name string) {
	o.ShortName = name
}