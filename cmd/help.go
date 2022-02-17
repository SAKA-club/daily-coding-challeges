package cmd

type Help struct{}

func NewHelp() *Help {
	return &Help{}
}

func (h Help) Name() string {
	return "help"
}

func (h Help) Description() string {
	return "prints allowed commands"
}

func (h Help) Options() map[string]string {
	return nil
}

func (h Help) Invoke(args []string) (bool, error) {
	return true, nil
}
