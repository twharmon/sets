# Sets
Generic sets for Go.

![](https://github.com/twharmon/sets/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/sets)](https://goreportcard.com/report/github.com/twharmon/sets) [![codecov](https://codecov.io/gh/twharmon/sets/branch/main/graph/badge.svg?token=K0P59TPRAL)](https://codecov.io/gh/twharmon/sets)

## Documentation
For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/sets).

## Install
`go get github.com/twharmon/sets`

## Usage
```go
package main

import (
	"fmt"

	"github.com/twharmon/sets"
)

func main() {
	// Create a new set of ints
	s := sets.New[int]()

	// Add some values
	s.Add(1, 2, 3)

	// Check for a value
	s.Contains(2) // true

	// Remove a value
	s.Remove(2)

	// Get slice of all values in set
	s.Slice() // [1 3]

	// Merge another set into s
	s.Merge(sets.New(2, 4))

	// Clear the set
	s.Clear()
}
```

## Contribute
Make a pull request.
