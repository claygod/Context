# Context
Context library stores designed to store the application context.

[![API documentation](https://godoc.org/github.com/claygod/Context?status.svg)](https://godoc.org/github.com/claygod/Context)
[![Go Report Card](https://goreportcard.com/badge/github.com/claygod/Context)](https://goreportcard.com/report/github.com/claygod/Context)

# Usage

An example of using the Context Library:

```go
package main

import (
	"fmt"
	"github.com/claygod/Context"
)

func main() {
	fmt.Print("\nTesting Context library")

	fmt.Print("\n1) create and fill the root context:")
	c1 := Context.New()
	c1.Set("a", 5)
	c1.Set("b", 6)
	fmt.Print("\n Read variable in c1: 'a' -> ", c1.Get("a")) // --> 5
	fmt.Print("\n Read variable in c1: 'b' -> ", c1.Get("a")) // --> 6

	fmt.Print("\n2) create a context branch and fill the variables with the same key but different valuest:")
	c2 := c1.Fix()
	c2.Set("a", 15)
	c2.Set("b", 16)
	fmt.Print("\n Read variable in c2: 'a' -> ", c2.Get("a")) // --> 15
	fmt.Print("\n Read variable in c2: 'b' -> ", c2.Get("a")) // --> 16

	fmt.Print("\n3) creating more context branch and add one variable with the same key but different value:")
	c3 := c1.Fix()
	c3.Set("a", 25)
	fmt.Print("\n Read variable in c3: 'a' -> ", c3.Get("a")) // --> 25
	fmt.Print("\n Read variable in c3: 'b' -> ", c3.Get("b")) // --> 6 !!!
}

```

# API

Methods:
-  *New* - create a new Context
-  *Set* - add the variable in context.
-  *Get* - get variable.
-  *Fix* - fix the state to create a branch.

