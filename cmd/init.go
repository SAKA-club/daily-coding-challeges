package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type Init struct {
	username string
}

func NewInit(username string) *Init {
	return &Init{
		username: username,
	}
}

func (i Init) Name() string {
	return "init"
}

func (i Init) Description() string {
	return "creates plugin for your solution"
}

func (i Init) Options() map[string]string {
	return map[string]string{
		"[-d]": "dd-MM-YYYY to specify specific day ",
	}
}

func (i Init) Invoke(args []string) (bool, error) {
	if i.username == "" {
		return false, errors.New(fmt.Sprintf("username required"))
	}

	date, err := ParseDate(args[1:])
	if err != nil {
		return false, err
	}

	if !ProblemExists(*date) {
		return false, errors.New(fmt.Sprintf("invalid date provided: %v", *date))
	}

	problemDir := path.Join("problems", strconv.Itoa(date.Year()), strings.ToLower(date.Month().String()), fmt.Sprintf("%02d", date.Day()))
	solutionPath := path.Join(problemDir, "solutions", i.username+".go")
	_, err = ioutil.ReadFile(solutionPath)
	if err == nil {
		return false, errors.New(fmt.Sprintf("solution file already exists: %s", solutionPath))
	}

	err = createTemplate(problemDir, solutionPath)

	return false, err
}

func createTemplate(problemDir string, solution string) error {
	templateData, err := ioutil.ReadFile(path.Join(problemDir, "template.txt"))
	if err != nil {
		return errors.New(fmt.Sprintf("template file does not exist: %s", path.Join(problemDir, "template.txt")))
	}

	ut, err := template.New(solution).Parse(string(templateData))
	if err != nil {
		return errors.New(fmt.Sprintf("error parsing template: %s", err.Error()))
	}

	solutionDir := path.Join(problemDir, "solutions")
	if _, err = os.Stat(solutionDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(solutionDir, 0755)
		}
	}
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create or access solutions directory: %s", solutionDir))
	}

	f, err := os.Create(solution)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating solution file: %s", err.Error()))
	}
	defer f.Close()

	buffer := bufio.NewWriter(f)
	err = ut.Execute(buffer, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("error writing to solution file: %s", err.Error()))
	}

	// flush buffered data to the file
	if err = buffer.Flush(); err != nil {
		return errors.New(fmt.Sprintf("error flushing : %s", err.Error()))
	}

	println(fmt.Sprintf("Created solution file: %s", solution))
	return nil
}
