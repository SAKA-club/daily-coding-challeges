package main

import (
	"club.saka/daily-coding-challeges/cmd"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	title, _ := ioutil.ReadFile("copy/daily_coding_challenge_ascii_art.txt")
	fmt.Println(string(title) + "\n")

	cmdOpts := map[string]cmd.CmdOption{
		"exit":     cmd.NewExit(),
		"init":     nil,
		"list":     nil,
		"validate": nil,
		"stats":    nil,
	}
	helpCmd := cmd.NewHelp(cmdOpts)
	cmdOpts["help"] = helpCmd

	execute(cmdOpts)
}

func execute(cmdOpts map[string]cmd.CmdOption) {
	var err error
	var rawCmds string
	var cmds []string
	exit := false

	for err == nil && !exit {
		print("\n$>")
		_, err = fmt.Scanln(&rawCmds)
		if err != nil {
			return
		}

		cmds = strings.Split(rawCmds, " ")

		// Check if argument is an allowed command
		cmdAllowed := false
		if len(cmds) >= 1 {
			_, cmdAllowed = cmdOpts[cmds[0]]
		}

		// Print help if there is no argument, argument is help, or cmd is not allowed
		if len(cmds) < 1 || cmds[0] == "help" || !cmdAllowed {
			exit, err = cmdOpts["help"].Invoke(cmds)
		} else {
			if cmdOpts[cmds[0]] == nil {
				println("\nnot implemented\n")
				exit, err = cmdOpts["help"].Invoke(cmds)
			} else {
				exit, err = cmdOpts[cmds[0]].Invoke(cmds)
			}
		}
	}

	if err != nil {
		println(err)
	}
}
