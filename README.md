# INTree for Go

Static, flat **IN**terval **Tree** implementation for reverse range searches.

The flat tree structure using Slices makes traversal very efficient, very fast, and with almost no memory footprint other than the range limits.

Further scientific reading about the adapted algorithm and comparisons between different approaches (in C/C++) can be found [here](https://github.com/lh3/cgranges).


## Behaviour:

* INTree will build the tree once (**static; no updates after creation**)
* INTree returns indices to the initial `[]Bounds` array

## Usage:

INTree currently supports finding all bounds for a simple float value.

### Import
```
import (
    "github.com/geozelot/intree"
)
```

### Example usage

```go
package main

import (
    "github.com/geozelot/intree"
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
