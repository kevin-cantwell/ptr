## Overview

This package enables in-line addressing of literals and other unassigned values by providing generic functions for referencing values or de-referencing pointers.

```go
var p string = "foo"
fmt.Printf("Pointer value of p: %v\n", ptr.P(p))
fmt.Printf("Deref value of *p: %v\n", ptr.V(&p))

var i *int64 = nil
fmt.Printf("Pointer value of i: %v\n", ptr.P(i))
fmt.Printf("Pointer value of *i: %v\n", ptr.V(&i))
```