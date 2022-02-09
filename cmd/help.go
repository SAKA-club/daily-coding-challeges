package cmd

import (
	"fmt"
	"os"
	"sort"
)

type CmdOption interface {
	Name() string
	Description() string
	Invoke([]string)
}

type HelpCmd struct {
	cmdOpts map[string]CmdOption
}

func NewHelpCmd(opts map[string]CmdOption) *HelpCmd {
	return &HelpCmd{
		cmdOpts: opts,
	}
}

func (h HelpCmd) Name() string {
	return "help"
}

func (h HelpCmd) Description() string {
	return "prints all available commands"
}

func (h HelpCmd) Invoke(args []string) {
	var err string
	if len(args) <= 1 {
		err = "no command provided"
	} else if cmd, allowed := h.cmdOpts[os.Args[1]]; !allowed {
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
}
