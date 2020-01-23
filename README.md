## INTree
#### Fast static Interval Tree for Go

Static, flat Interval Tree implementation for reverse range searches (**which intervals include a given value**). The tree is realized using Go Slices only; **memory usage is low, performance extremely high**!

___

#### Behaviour:

* INTree will build the tree once (**static; no updates after creation**)
* INTree returns indices to the initial `[]Bounds`!

___

#### Usage:

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
