package cmd

import (
	"club.saka/daily-coding-challeges/common"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"plugin"
	"strings"
)

const testInputPath = "test_inputs.json"

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

	newFn, err := testPlugin.Lookup("New")
	if err != nil {
		return errors.New(fmt.Sprintf("test plugin function `New` could not be found: %s", err.Error()))
	}

	pwd := path.Join(basePath, testInputPath)
	logger := log.With().Str("pwd", pwd).Logger()

	problem := newFn.(func(pwd string, logger *zerolog.Logger) *common.Problem)(pwd, &logger)
	if problem == nil {
		return err
	}

	harness, err := (*problem).RunTests()
	logger.Info().Int("passed", harness.Passed).Int("failed", harness.Failed).Dur("duration", harness.End.Sub(harness.Start))

	return nil
}
