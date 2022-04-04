package common

type Problem interface {
	RunTests() (*TestHarness, error)
}
