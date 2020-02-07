# INTree for Go

Very fast static, flat **IN**terval **Tree** implementation for reverse range searches.

Highly efficient and with almost no memory footprint other than the stored ranges.

Further scientific reading about the adapted algorithm and comparisons between different approaches (in C/C++) can be found [here](https://github.com/lh3/cgranges).


# Behaviour

* INTree will build the tree once (**static; no updates after creation**)
* INTree returns indices to the initial `[]Bounds` array
* INTree currently supports finding all bounds for a single `float64` value

# Usage

## API ([GoDoc](https://godoc.org/github.com/geozelot/intree))

### `type Bounds`

`Bounds{}` is the main interface expected by `NewINTree()`; requires `Limits()` method to access interval limits.

```go
type Bounds interface {
    Limits() (Lower, Upper float64)
}
```

### `type INTree`

`INTree{}` is the main package object; holds Slice of reference indices and the respective interval limits.
https://play.golang.org/p/rrV_duXi7eh
```go
type INTree struct {
    // contains filtered or unexported fields
}
```

### `func NewINTree`

`NewINTree()` is the main initialization function; creates the tree from the given Slice of Bounds.

```go
func NewINTree(bnds []Bounds) *INTree
```

### `func (*INTree) Including`

`Including()` is the main entry point for bounds searches; traverses the tree and collects intervals that overlap with the given value.

```go
func (inT *INTree) Including(val float64) []int
```

## Import
```go
import (
    "github.com/geozelot/intree"
)
```

## Examples

#### Simple `Bounds{}` interface implementation:

```go
// SimpleBounds is a simple Struct implicitly implementing the Bounds interface.
type SimpleBounds struct {

  Lower, Upper float64

}

// Limits accesses the interval limits.
func (sb *SimpleBounds) Limits() (float64, float64) {

  return sb.Lower, sb.Upper

}
```

#### Test Setup:

```go
package main

import (

    "github.com/geozelot/intree"
    "fmt"

)

// defining simple Struct holding interval limits
type SimpleBounds struct {
  Lower, Upper float64
}

  // add method to access limits; implicitly implements INTree.Bounds interface
  func (sb *SimpleBounds) Limits() (float64, float64) {
    return sb.Lower, sb.Upper
  }

func main() {

  // create typed var
  var tree *intree.INTree
  
  // create example bounds
  inputBounds := []intree.Bounds{

    &intree.SimpleBounds{Lower: 4.0, Upper: 6.0},   // match
    &intree.SimpleBounds{Lower: 5.0, Upper: 7.0},
    &intree.SimpleBounds{Lower: 4.0, Upper: 8.0},   // match
    &intree.SimpleBounds{Lower: 1.0, Upper: 3.0},
    &intree.SimpleBounds{Lower: 7.0, Upper: 9.0},
    &intree.SimpleBounds{Lower: 3.0, Upper: 6.0},   // match
    &intree.SimpleBounds{Lower: 2.0, Upper: 3.0},
    &intree.SimpleBounds{Lower: 5.3, Upper: 7.9},
    &intree.SimpleBounds{Lower: 3.2, Upper: 7.5},   // match
    &intree.SimpleBounds{Lower: 4.4, Upper: 5.1},
    &intree.SimpleBounds{Lower: 4.1, Upper: 4.9},   // match
    &intree.SimpleBounds{Lower: 4.1, Upper: 4.9},   // match, same interval
    &intree.SimpleBounds{Lower: 1.3, Upper: 3.1},
    &intree.SimpleBounds{Lower: 7.9, Upper: 8.9},

  }

  // initialize new INTree and create tree from inputBounds
  tree = intree.NewINTree(inputBounds)

  // parse return Slice with indices referencing inputBounds
  for _, matchedIndex := range tree.Including(4.3) {

    // using INTree.Bounds interface method to access limits
    lowerLimit, upperLimit := inputBounds[matchedIndex].Limits()

    fmt.Printf("Match at inputBounds index %2d with range [%.1f, %.1f]\n", matchedIndex, lowerLimit, upperLimit)

    /*
      Match at inputBounds index 11 with range [4.1, 4.9]
      Match at inputBounds index 10 with range [4.1, 4.9]
      Match at inputBounds index  5 with range [3.0, 6.0]
      Match at inputBounds index  2 with range [4.0, 8.0]
      Match at inputBounds index  0 with range [4.0, 6.0]
      Match at inputBounds index  9 with range [3.2, 7.5]
    */

  }

}
```

#### Try on [Go Playground](https://play.golang.org/p/rrV_duXi7eh).

____

##### Inspired by this great [KDTree implementation](https://github.com/mourner/kdbush) for JavaScript and adapted from this excellent [Go port](https://github.com/MadAppGang/kdbush).
