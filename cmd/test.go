package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"plugin"
	"strings"
	"time"
)

type Test struct {
	username string
}

func NewTest(username string) *Test {
	return &Test{
		username: username,
	}
}

func (t Test) Name() string {
	return "test"
}

func (t Test) Description() string {
	return "runs tests against your solution"
}

func (t Test) Options() map[string]string {
	return map[string]string{
		"[-d]": "dd-MM-YYYY to specify specific day ",
	}
}

func (t Test) Invoke(args []string) (bool, error) {
	if t.username == "" {
		return false, errors.New(fmt.Sprintf("username required"))
	}

	date, err := ParseDate(args[1:])
	if err != nil {
		return false, err
	}

	if !ProblemExists(*date) {
		return false, errors.New(fmt.Sprintf("invalid date provided: %v", *date))
	}

	basePath := fmt.Sprintf("problems/%d/%s/%02d", date.Year(), strings.ToLower(date.Month().String()), date.Day())

	err = runTests(basePath, t.username)

	return false, err
}

func runTests(basePath string, username string) error {
	solutionPath := fmt.Sprintf("%s/solutions/%s.go", basePath, username)
	_, err := ioutil.ReadFile(solutionPath)
	if err != nil {
		return errors.New(fmt.Sprintf("solution file does not exists: %s", solutionPath))
	}

	println(fmt.Sprintf("Running tests for %s", solutionPath))

	// Build the test plugin if it hasn't been built yet
	testPluginPath := basePath + "/test.so"
	_, err = ioutil.ReadFile(testPluginPath)
	if err != nil {
		for _, s := range os.Environ() {
			println(s)
		}

		// TODO: Add additional args based on XPC_SERVICE_NAME
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", testPluginPath, basePath+"/test.go")
		if err = cmd.Run(); err != nil {
			return errors.New(fmt.Sprintf("test plugin could not be compiled: %s", err.Error()))
		}

		_, err = ioutil.ReadFile(testPluginPath)
		if err != nil {
			return errors.New(fmt.Sprintf("test plugin could not be found: %s", err.Error()))
		}
	}

	testPlugin, err := plugin.Open(testPluginPath)
	if err != nil {
		return errors.New(fmt.Sprintf("test plugin could not be opened: %s", err.Error()))
	}

	testCount, err := testPlugin.Lookup("TestCount")
	if err != nil {
		return errors.New(fmt.Sprintf("test plugin constant TestCount could not be found: %s", err.Error()))
	}

	testRunner, err := testPlugin.Lookup("TestRunner")
	if err != nil {
		return errors.New(fmt.Sprintf("test plugin constant TestRunner could not be found: %s", err.Error()))
	}

	start := time.Now()
	println(fmt.Sprintf("Executing %v tests...", *testCount.(*int)))
	testRunner.(func())()
	println(fmt.Sprintf("Execution time: %v seconds", time.Now().Sub(start).Seconds()))

	return nil
}
