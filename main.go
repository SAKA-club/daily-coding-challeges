package main

import (
	"bufio"
	"club.saka/daily-coding-challeges/cmd"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Command interface {
	Name() string
	Description() string
	Invoke([]string) (bool, error)
	Options() map[string]string
}

func main() {
	title, _ := ioutil.ReadFile("copy/daily_coding_challenge_ascii_art.txt")
	fmt.Println(string(title) + "\n")

	cmds := map[string]Command{
		"exit":     cmd.NewExit(),
		"help":     cmd.NewHelp(),
		"init":     nil,
		"list":     cmd.NewList(),
		"validate": nil,
		"stats":    nil,
	}

	execute(cmds)
}

func execute(cmds map[string]Command) {
	var err error
	var args []string
	exit := false

	for !exit {
		print("\n$>")

		// Get arguments
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			args = strings.Split(scanner.Text(), " ")
		}

		// Check if first argument is an allowed command
		isValidCmd := false
		if len(args) >= 1 {
			_, isValidCmd = cmds[args[0]]
		}

		// Print help if there is no argument, command is help, or command is invalid
		if len(args) < 1 || args[0] == "help" || !isValidCmd {
			exit, err = help(cmds, args)
		} else {
			if cmds[args[0]] == nil {
				fmt.Printf("\n%s is not implemented\n", args[0])
				exit, err = help(cmds, args)
			} else {
				exit, err = cmds[args[0]].Invoke(args)
			}
		}

		if err != nil {
			println(err.Error())
		}
	}
}

func help(commands map[string]Command, args []string) (bool, error) {
	var err string
	if len(args) < 1 {
		err = "no command provided"
	} else if cmd, allowed := commands[args[0]]; !allowed {
		err = fmt.Sprintf("%v: no command found", cmd)
	}

	if err != "" {
		println(fmt.Sprintf("\n%v\n", err))
	}

	cmds := make([]string, 0, len(commands))
	for cmd := range commands {
		cmds = append(cmds, cmd)
	}
	sort.Strings(cmds)

	println("Available commands:")
	for _, cmd := range cmds {
		if commands[cmd] == nil {
			println(fmt.Sprintf("    %-10s: %v", cmd, "not implemented"))
			continue
		}

		println(fmt.Sprintf("    %-10s: %v", cmd, commands[cmd].Description()))
		if opts := commands[cmd].Options(); opts != nil {
			for option, desc := range opts {
				println(fmt.Sprintf("                %s %s", option, desc))
			}
		}
	}

	return false, nil
}
