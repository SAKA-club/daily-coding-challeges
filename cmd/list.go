package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
)

const TimeLayout = "02-01-2006"

type List struct{}

func NewList() *List {
	return &List{}
}

func (l List) Name() string {
	return "list"
}

func (l List) Description() string {
	return "Prints out the prompt for the challenge of the day"
}

func (l List) Options() map[string]string {
	return map[string]string{
		"[-d]": "dd-MM-YYYY to specify specific day ",
	}
}

func (l List) Invoke(args []string) (exit bool, err error) {
	// remove first arg (list) from args list
	args = args[1:]
	date := time.Now()

	if len(args) > 0 {
		if len(args) != 2 {
			return false, errors.New(fmt.Sprintf("invalid number of arguments for `list`: %d", len(args)))
		}

		if args[0] != "-d" {
			return false, errors.New(fmt.Sprintf("invalid option for `list` command: %s", args[0]))
		}

		date, err = time.Parse(TimeLayout, args[1])
		if err != nil {
			return false, errors.New(fmt.Sprintf("invalid date: %s", args[1]))
		}
	}

	path := fmt.Sprintf("problems/%d/%s/%02d/README.md", date.Year(), strings.ToLower(date.Month().String()), date.Day())
	_, err = ioutil.ReadFile(path)
	if err != nil {
		return false, errors.New(noProbErr(date, path))
	} else {
		readme, err := render(path)
		if err != nil {
			return false, err
		}

		println(fmt.Sprintf("%s", readme))
	}

	return false, nil
}

func noProbErr(date time.Time, path string) string {
	formattedDate := date.Format(TimeLayout)
	return fmt.Sprintf("The problem for %s isn't posted yet or your local repo is out of date:\n\tPath: %s", formattedDate, path)
}

func render(path string) ([]byte, error) {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to read file at %s", path))
	}

	return markdown.Render(string(source), 80, 6), nil
}
