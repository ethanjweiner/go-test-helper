package testhelper

import (
	"testing"
)

// Include desired assertions in exports

type TestCase[Input any] struct {
	Name     string
	Args     Input
	Expected any
	Test     func(TestCase[Input])
}

func Run_Tests[Input any](t *testing.T, tests []TestCase[Input], defaultTest func(TestCase[Input])) {
	for _, test := range tests {
		t.Run(test.Name, func(*testing.T) {
			if test.Test == nil {
				defaultTest(test)
			} else {
				test.Test(test)
			}
		})
	}
}
