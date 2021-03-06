package main

import (
	"club.saka/daily-coding-challeges/common"
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

// StringCompression perform basic string compression using the counts of repeated characters
func StringCompression(s string) string {
	prev, builder, count := "", "", 1
	for i, v := range s {
		// Add to builder
		if v := string(v); prev != v || i == len(s)-1 {
			builder += prev
			if count > 1 {
				builder += strconv.Itoa(count)
			}
			prev, count = v, 1
		} else {
			count++
		}

		// Handle last case scenario
		if i == len(s)-1 {
			builder += prev
			if count > 1 {
				builder += strconv.Itoa(count)
			}
		}
	}

	return builder
}

/* DO NOT EDIT THE FOLLOWING */

type TestInput struct {
	S string `json:"s"`
}

type Test struct {
	Inputs  map[string]TestInput `json:"inputs"`
	Outputs map[string]string    `json:"outputs"`
}

type Problem struct {
	tests  *map[string]Test
	logger *zerolog.Logger
}

func New(pwd string, logger *zerolog.Logger) *common.Problem {
	tests := map[string]Test{}
	err := common.ImportTestData(pwd, &tests)
	if tests == nil || err != nil {
		log.Log().Err(err).Msg("no test data")
	}

	var p common.Problem
	p = Problem{
		tests:  &tests,
		logger: logger,
	}

	return &p
}

func (p Problem) RunTests() (*common.TestHarness, error) {
	if p.tests == nil {
		return nil, errors.New("no test data")
	}

	harness := common.TestHarness{
		Total: len(*p.tests),
	}

	harness.Start = time.Now()
	for title, t := range *p.tests {
		for header, in := range t.Inputs {
			ll := p.logger.With().Str("test", title).Str("subtest", header).Interface("input", in).Logger()
			expected := t.Outputs[header]
			output := StringCompression(in.S)

			if expected != output {
				harness.Failed += 1
				ll.Error().Interface("expected_output", expected).Interface("output", output).Msg("output mismatch")
				break
			}

			ll.Debug().Interface("output", output).Msg("Success")
			harness.Passed += 1
		}
	}

	harness.End = time.Now()
	return &harness, nil
}
