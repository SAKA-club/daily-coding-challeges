package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func ProblemExists(date time.Time) bool {
	path := fmt.Sprintf("problems/%d/%s/%02d/README.md", date.Year(), strings.ToLower(date.Month().String()), date.Day())
	_, err := ioutil.ReadFile(path)
	return err == nil
}

func ParseDate(args []string) (*time.Time, error) {
	date := time.Now()

	if len(args) > 0 {
		if len(args) != 2 {
			return nil, errors.New(fmt.Sprintf("invalid number of arguments for `list`: %d", len(args)))
		}

		if args[0] != "-d" {
			return nil, errors.New(fmt.Sprintf("date operation must be [-d], not: %s", args[0]))
		}

		var err error
		date, err = time.Parse(TimeLayout, args[1])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("invalid date: %s", args[1]))
		}
	}

	return &date, nil
}
