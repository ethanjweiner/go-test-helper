## Using go-test-helper

```go
package add

import (
	"testing"

	"github.com/ethanjweiner/go-test-helper/testhelper"
	"github.com/stretchr/testify/assert" // Your library of choice
)

func Add(a int, b int) int {
	return a + b
}

func Test_Add(t *testing.T) {
	// Define the types of your input
	type input struct {
		a int
		b int
	}

	// Define your test cases
	tests := []testhelper.TestCase[input]{
		{
			Name:     "1 + 2",
			Args:     input{1, 2},
			Expected: 3,
		},
		{
			Name:     "1 + -1",
			Args:     input{1, -1},
			Expected: 0,

			// Optional: Define custom testing logic for the test
			Test: func(test testhelper.TestCase[input]) {
				actual := Add(test.Args.a, test.Args.b)
				assert.NotNil(t, actual)
				assert.Equal(t, actual, test.Expected)
			},
		},
	}

	// Run all tests
	testhelper.Run_Tests(t, tests, func(test testhelper.TestCase[input]) {
		// Define the default testing logic here
		actual := Add(test.Args.a, test.Args.b)
		assert.Equal(t, actual, test.Expected)
	})
}

```

If you wish, create a snippet of the above in your code editor for easier usage.