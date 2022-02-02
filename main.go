package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	title, _ := ioutil.ReadFile("copy/daily_coding_challenge_ascii_art.txt")
	fmt.Println(string(title) + "\n\n\n")

	// Options are:
	//  help(h): Prints out the allowed arguments
	//  list(l): Prints out the prompt for the challenge of the day
	//  init(i): Initializes the file for your challenge of the day
	//  test(t) [all, day]: Tests your code for the challenge of today (or all/specific day if specified)

	helpcmd := flag.NewFlagSet("help", flag.ExitOnError)
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

	testcmd := flag.NewFlagSet("test", flag.ExitOnError)
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
	case "test":
		testcmd.Parse(os.Args[2:])
		fmt.Println("Test is not yet implemented")
	default:
		fmt.Println("Help is not yet implemented")
		os.Exit(1)
	}
}
