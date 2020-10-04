package cli

type Option struct {
	Name 		string
	ShortName 	string
	ArgsType	[]string
	Description string
}

func (a *Action) AddOption(name, description string, argsType ...string) *Option {
	o := new(Option)
	o.Name = name
	o.Description = description
	o.ArgsType = argsType
	a.Options = append(a.Options, o)
	return o
}

func (o *Option) AddShortName(name string) {
	o.ShortName = name
}