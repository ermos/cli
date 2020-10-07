package cli

type Option struct {
	Name 		string
	ShortName 	string
	ArgsType	[]string
	Description string
}

type OptionValue struct {
	String 		[]string
	Interface 	[]interface{}
}

func AddOption(name, description string, argsType ...string) *Option {
	o := addOption(name, description, argsType...)
	c.options = append(c.options, o)
	return o
}

func (a *Action) AddOption(name, description string, argsType ...string) *Option {
	o := addOption(name, description, argsType...)
	a.options = append(a.options, o)
	return o
}

func addOption (name, description string, argsType ...string) *Option {
	o := new(Option)
	o.Name = name
	o.Description = description
	for _, arg := range argsType {
		o.ArgsType = append(o.ArgsType, arg)
	}
	return o
}

func (o *Option) AddShortName(name string) {
	o.ShortName = name
}