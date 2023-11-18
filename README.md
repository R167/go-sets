# Go Sets!
A minimal set library for golang. This tinylib eschews any dependencies and implements common, optimized set operations on `type Set[E comparable] map[E]struct{}`

Supported set operations:
- Union: A⋃B
- Intersection: A⋂B
- Difference: A-B

Supported checks:
- Subset: A⊆B
- Proper subset: A⊂B
- Superset: A⊇B
- Proper superset: A⊃B
- Equality: A=B

## Usage
```go
package main

import (
  "fmt"

  "github.com/R167/go-sets"
)

func main() {
  a := sets.NewSet(1, 2, 3)
  b := sets.NewSet(3, 4, 5)

  fmt.Println(a.Union(b)) // {1, 2, 3, 4, 5}
  fmt.Println(a.Intersection(b)) // {3}
  fmt.Println(a.Difference(b)) // {1, 2}
  fmt.Println(a.Subset(b)) // false
  fmt.Println(a.ProperSubset(b)) // false
  fmt.Println(a.Superset(b)) // false
  fmt.Println(a.ProperSuperset(b)) // false
  fmt.Println(a.Equality(b)) // false
}
```
