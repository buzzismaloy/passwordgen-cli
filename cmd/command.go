package cmd

type Command struct {
	Use   string
	Short string
	Long  string
	Run   func()

	Subcommands []*Command
}

func (c *Command) Execute() error {
	if c.Run != nil {
		c.Run()
		return nil
	}

	for _, sc := range c.Subcommands {
		sc.Execute()
	}

	return nil
}
