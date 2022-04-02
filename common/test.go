package common

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"time"
)

type TestRunner interface {
	RunTests(ctx context.Context, tests interface{}) (TestHarness, error)
}

type TestHarness struct {
	Total  int
	Passed int
	Failed int
	Start  time.Time
	End    time.Time
}

func ImportTestData(pwd string, tests interface{}) error {
	testData, err := ioutil.ReadFile(pwd)
	if err != nil {
		return err
	}

	err = json.Unmarshal(testData, &tests)
	if err != nil {
		return err
	}

	return nil
}
