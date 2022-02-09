package cmd

type Exit struct {
	cmdOpts map[string]CmdOption
}

func NewExit() *Exit {
	return &Exit{}
}

func (e Exit) Name() string {
	return "exit"
}

func (e Exit) Description() string {
	return "quits the program"
}

func (e Exit) Invoke(args []string) (bool, error) {
	return true, nil
}
