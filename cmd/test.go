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
	solutionPath := path.Join(basePath, "solutions", username+".go")
	pluginPath := path.Join(basePath, "solutions", fmt.Sprintf("%s_test.so", username))
	_, err := ioutil.ReadFile(solutionPath)
	if err != nil {
		return errors.New(fmt.Sprintf("solution file does not exists: %s", solutionPath))
	}

	// Remove old plugin
	err = os.Remove(pluginPath)
	if !os.IsNotExist(err) {
		return err
	}

	// TODO - BUG
	// The problem as I see it is that we are using the command line to build the plugin and then trying to run the
	// plugin from inside this program. We can't guarantee that these will work unless we can freeze the build falgs.
	// This works if we just run the program, but will fail if we try to debug the program.
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginPath, solutionPath)
	if err = cmd.Run(); err != nil {
		return errors.New(fmt.Sprintf("test plugin could not be compiled: %s", err.Error()))
	}
	_, err = ioutil.ReadFile(pluginPath)
	if err != nil {
		return errors.New(fmt.Sprintf("test plugin could not be found: %s", err.Error()))
	}

	testPlugin, err := plugin.Open(pluginPath)
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
		return errors.New("test plugin function could not be instantiated")
	}

	harness, err := (*problem).RunTests()
	logger.Info().Int("passed", harness.Passed).Int("failed", harness.Failed).Dur("duration", harness.End.Sub(harness.Start)).Send()

	return nil
}
