package main

import (
	"club.saka/daily-coding-challeges/cmd"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	title, _ := ioutil.ReadFile("copy/daily_coding_challenge_ascii_art.txt")
	fmt.Println(string(title) + "\n")

	cmds := map[string]cmd.CmdOption{
		"init":     nil,
		"list":     nil,
		"validate": nil,
	}
	helpCmd := cmd.NewHelpCmd(cmds)
	cmds["help"] = helpCmd

	// Check if argument is an allowed command
	cmdAllowed := false
	if len(os.Args) > 1 {
		_, cmdAllowed = cmds[os.Args[1]]
	}

	// Print help if there is no argument, argument is help, or cmd is not allowed
	if len(os.Args) <= 1 || os.Args[1] == "help" || !cmdAllowed {
		cmds["help"].Invoke(os.Args)
		os.Exit(1)
	}

	cmds[os.Args[1]].Invoke(os.Args[2:])

	// Options are:
	//
	//  list: Prints out the prompt for the challenge of the day
	//  init: Initializes the file for your challenge of the day
	//  validate [all, day]: Tests your code for the challenge of today (or all/specific day if specified)

	helpcmd := flag.NewFlagSet("help", flag.ContinueOnError)
	// TODO:
	// print all the options

	listcmd := flag.NewFlagSet("list", flag.ExitOnError)
	// TODO:
	// completed := lists all completed challenges similar to the GitHub days with a commit
	// day := lists the challenge for that specified day

	initcmd := flag.NewFlagSet("list", flag.ExitOnError)
	// TODO:
	// day := inits the solution file for that specified day
	// [arg] username := override the env variable for username

	validatecmd := flag.NewFlagSet("validate", flag.ExitOnError)
	// TODO:
	// day := runs the tests for that given day

	switch os.Args[1] {
	case "help":
		helpcmd.Parse(os.Args[2:])
		fmt.Println("Help is not yet implemented")
	case "list":
		listcmd.Parse(os.Args[2:])
		fmt.Println("List is not yet implemented")
	case "init":
		initcmd.Parse(os.Args[2:])
		fmt.Println("Init is not yet implemented")
	case "validate":
		validatecmd.Parse(os.Args[2:])
		fmt.Println("Test is not yet implemented")
	default:
		fmt.Println("Help is not yet implemented")
	}
}
