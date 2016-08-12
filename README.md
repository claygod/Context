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
	ctx := Context.New()
	fmt.Print("\nTesting Context library:")
	fmt.Print("\n Add key 'abc' -> variable 'ok'")
	ctx.Set("abc", "ok")
	fmt.Print("\n Read variable: 'abc' -> ", ctx.Get("abc")) // ok
}

```

# API

Methods:
-  *New* - create a new Context
-  *Set* - add the variable in context.
-  *Get* - get variable.

