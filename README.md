## Using go-test-helper

```go
package mypackage

import (
	"testing"
	"github.com/stretchr/testify/assert"
  "github.com/ethanjweiner/go-test-helper"
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
	tests := []testkit.TestCase[input]{
		{
			Name:     "1 + 2",
			Args:     input{1, 2},
			Expected: 3,
		},
	}

  // Run all tests
	Run_Tests(t, tests, func(test testkit.TestCase[input]) {
    // Add assertings here with your preferred testing library
		actual := Add(test.Args.a, test.Args.b)
		assert.Equal(t, actual, test.Expected)
	})
}
```

If you wish, create a snippet of the above in your code editor for easier usage.