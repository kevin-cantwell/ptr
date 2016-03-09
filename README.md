## Overview

This package enables in-line addressing of literals and other unassigned values by providing mapping functions for each predeclared identifier that accept a value and return its address.

----------------

The Go programming language does not allow you to take the address of an unassigned value. I.e., this is not possible:

```go
var foo *string = &"bar" // cannot take the address of "bar"
```
Try it: http://play.golang.org/p/mnmeFzQ0SW



The solution is generally to assign "bar" to a variable, then take its address:

```go
bar := "bar"
var foo *string = *bar // Totally cool
```
Try it: http://play.golang.org/p/IO12_CTwkj



Another way to solve this problem is to return a pointer value directly from a function:

```go
func bar(s string) *string {
  return &s
}
// ...
var foo *string = bar("bar")
```
Try it: http://play.golang.org/p/OUcCdBJUcb

This package takes the last approach and makes it available as helper functions.


### Example

```go
package main

import "github.com/kevin-cantwell/p"

type MyGodItsFullOfStars struct {
  A *int
  B *string
  C *float64
}

func main() {
  stars := MyGodItsFullOfStars{
    A: p.Int(123),
    B: p.String("foo"),
    C: p.Float64(33.3),
  }
  fmt.Println(stars)
}
```