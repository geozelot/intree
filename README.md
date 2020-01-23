# INTree for Go

Static, flat Interval Tree implementation for reverse range searches (**which intervals include a given value**).

The flat tree structure using Go Slices makes traversal very fast, with almost no memory footprint other than the stored ranges.

**Stil testing; handle with care!**

## Behaviour:

* INTree will build the tree once (**static; no updates after creation**)!
* INTree returns indices to the initial `[]Bounds`!
* INTree does currently not handle duplicate ranges correctly!

## Usage:

Currently the only supported query is to find all bounds for a simple value.

```go
// create dummy bounds
size := 100
bounds := make([]intree.Bounds, size, size)
rand.Seed(time.Now().UnixNano())
for i := 0; i < size; i++ {
    bounds[i] = &intree.Bounds{ Lower: i + rand.Intn(100), Upper: i + i + rand.Intn(100) }
}

// declare type
var tree *intree.INTree

// initialize new tree
tree = intree.NewINTree(bounds)

// find all nodes (bounds) that include the given value
for idx := range tree.Including(42) {
  fmt.Println("Found: ", bounds[idx])
}
```
____

##### Inspired by this great [KDTree implementation](https://github.com/mourner/kdbush) for JavaScript and adapted from this excellent [Go port](https://github.com/MadAppGang/kdbush).
