package testhelper

import (
	"testing"
)

// Include desired assertions in exports

type TestCase[Input any] struct {
	Name     string
	Args     Input
	Expected any
}

func Run_Tests[Input any](t *testing.T, tests []TestCase[Input], testFunc func(TestCase[Input])) {
	for _, test := range tests {
		t.Run(test.Name, func(*testing.T) { testFunc(test) })
	}
}
