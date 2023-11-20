# Go Sets!
A minimal set library for golang. This tinylib eschews any dependencies and implements common, optimized set operations on `type Set[E comparable] map[E]struct{}`

[![Go Reference](https://pkg.go.dev/badge/github.com/R167/go-sets.svg)](https://pkg.go.dev/github.com/R167/go-sets) [![Go Report Card](https://goreportcard.com/badge/github.com/R167/go-sets)](https://goreportcard.com/report/github.com/R167/go-sets) [![CI](https://github.com/R167/go-sets/actions/workflows/ci.yaml/badge.svg)](https://github.com/R167/go-sets/actions/workflows/ci.yaml)

Supported set operations:
- Union: A⋃B
- Intersection: A⋂B
- Difference: A-B

Supported checks:
- Subset: A⊆B
- Superset: A⊇B
- Equality: A=B

## Usage
```go
package main

import (
  "fmt"

  "github.com/R167/go-sets"
)

func main() {
  a := sets.New(1, 2, 3)
  b := sets.New(3, 4, 5)

  fmt.Println(a.Union(b)) // {1, 2, 3, 4, 5}
  fmt.Println(a.Intersect(b)) // {3}
  fmt.Println(a.Difference(b)) // {1, 2}

  a2 := sets.New(1, 2)
  fmt.Println(a2.Subset(a)) // true
  fmt.Println(a.Subset(a2)) // false
  fmt.Println(a.Equal(b)) // false
}
```
