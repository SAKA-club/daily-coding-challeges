package cmd

type Exit struct{}

func NewExit() *Exit {
	return &Exit{}
}

func (e Exit) Name() string {
	return "exit"
}

func (e Exit) Description() string {
	return "quits the program"
}

func (e Exit) Options() map[string]string {
	return nil
}

func (e Exit) Invoke(args []string) (bool, error) {
	return true, nil
}
