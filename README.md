# INTree for Go

Static, flat Interval Tree implementation for reverse range searches (**which intervals include a given value**).

The flat tree structure using Go Slices makes traversal very fast, with almost no memory footprint other than the stored ranges.

Current implementation is running on recusive traversal again; will replace with heap collecting loop ASAP.

**Stil testing; handle with caution!**

## Behaviour:

* INTree will build the tree once (**static; no updates after creation**)
* INTree returns indices to the initial `[]Bounds`
* INTree currently supports finding all bounds for a simple float value

## Usage:

### Import
```
import (
    "github.com/geozelot/intree-go"
)
```

### Example usage

```go
package main

import (
    "github.com/geozelot/intree-go"
    "fmt"
    "math/rand"
    "time"
)

func main() {
 
  // create dummy bounds
  size := 100
  bounds := make([]intree.Bounds, size, size)
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < size; i++ {
    bounds[i] = &intree.SimpleBounds{ Lower: float64(i + rand.Intn(100)), Upper: float64(i * 2 + rand.Intn(100))
  }

  // declare type
  var tree *intree.INTree

  // initialize new tree
  tree = intree.NewINTree(bounds)

  // find all nodes (bounds) that include the given value
  for _, idx := range tree.Including(float64(42)) {
    fmt.Println("Found: ", bounds[idx])
  }

}
```
____

##### Inspired by this great [KDTree implementation](https://github.com/mourner/kdbush) for JavaScript and adapted from this excellent [Go port](https://github.com/MadAppGang/kdbush).
