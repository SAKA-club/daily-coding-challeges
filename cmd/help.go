package cmd

import (
	"fmt"
	"sort"
)

type CmdOption interface {
	Name() string
	Description() string
	Invoke([]string) (bool, error)
}

type Help struct {
	cmdOpts map[string]CmdOption
}

func NewHelp(opts map[string]CmdOption) *Help {
	return &Help{
		cmdOpts: opts,
	}
}

func (h Help) Name() string {
	return "help"
}

func (h Help) Description() string {
	return "prints all available commands"
}

func (h Help) Invoke(args []string) (bool, error) {
	var err string
	if len(args) < 1 {
		err = "no command provided"
	} else if cmd, allowed := h.cmdOpts[args[0]]; !allowed {
		err = fmt.Sprintf("%v: no command found", cmd)
	}

	if err != "" {
		println(fmt.Sprintf("\n%v\n", err))
	}

	cmds := make([]string, 0, len(h.cmdOpts))
	for cmd := range h.cmdOpts {
		cmds = append(cmds, cmd)
	}
	sort.Strings(cmds)

	println("Available commands:")
	for _, cmd := range cmds {
		desc := "not implemented"
		if h.cmdOpts[cmd] != nil {
			desc = h.cmdOpts[cmd].Description()
		}
		println(fmt.Sprintf("    %-10s: %v", cmd, desc))
	}

	return false, nil
}
