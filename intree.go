// MIT License
//
// Copyright (c) 2022 geozelot (André Siefken)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package intree provides a very fast, static, flat, augmented interval tree for reverse range searches.
package intree

import (
	"math"
	"math/rand"
)

// Bounds is the main interface expected by NewINTree(); requires Limits method to access interval limits.
type Bounds interface {
	Limits() (Lower, Upper float64)
}

// INTree is the main package object;
// holds Slice of reference indices and the respective interval limits.
type INTree struct {
	idxs []int
	lmts []float64
}

// buildTree is the internal tree construction function;
// creates, sorts and augments nodes into Slices.
func (inT *INTree) buildTree(bnds []Bounds) {

	inT.idxs = make([]int, len(bnds))
	inT.lmts = make([]float64, 3*len(bnds))

	for i, v := range bnds {

		inT.idxs[i] = i
		l, u := v.Limits()

		inT.lmts[3*i] = l
		inT.lmts[3*i+1] = u
		inT.lmts[3*i+2] = 0

	}

	sort(inT.lmts, inT.idxs)
	augment(inT.lmts, inT.idxs)

}

// Including is the main entry point for bounds searches;
// traverses the tree and collects intervals that overlap with the given value.
func (inT *INTree) Including(val float64) []int {

	stk := []int{0, len(inT.idxs) - 1}
	res := []int{}

	for len(stk) > 0 {

		rb := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		lb := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if lb == rb+1 {
			continue
		}

		cn := int(math.Ceil(float64(lb+rb) / 2.0))
		nm := inT.lmts[3*cn+2]

		if val <= nm {

			stk = append(stk, lb)
			stk = append(stk, cn-1)

		}

		l := inT.lmts[3*cn]

		if l <= val {

			stk = append(stk, cn+1)
			stk = append(stk, rb)

			u := inT.lmts[3*cn+1]

			if val <= u {
				res = append(res, inT.idxs[cn])
			}

		}

	}

	return res

}

// NewINTree is the main initialization function;
// creates the tree from the given Slice of Bounds.
func NewINTree(bnds []Bounds) *INTree {

	inT := INTree{}
	inT.buildTree(bnds)

	return &inT

}

// augment is an internal utility function, adding maximum value of all child nodes to the current node.
func augment(lmts []float64, idxs []int) {

	if len(idxs) < 1 {
		return
	}

	max := 0.0

	for idx := range idxs {

		if lmts[3*idx+1] > max {
			max = lmts[3*idx+1]
		}

	}

	r := len(idxs) >> 1

	lmts[3*r+2] = max

	augment(lmts[:3*r], idxs[:r])
	augment(lmts[3*r+3:], idxs[r+1:])

}

// sort is an internal utility function, sorting the tree by lowest limits using Random Pivot QuickSearch
func sort(lmts []float64, idxs []int) {

	if len(idxs) < 2 {
		return
	}

	l, r := 0, len(idxs)-1

	p := rand.Int() % len(idxs)

	idxs[p], idxs[r] = idxs[r], idxs[p]
	lmts[3*p], lmts[3*p+1], lmts[3*p+2], lmts[3*r], lmts[3*r+1], lmts[3*r+2] = lmts[3*r], lmts[3*r+1], lmts[3*r+2], lmts[3*p], lmts[3*p+1], lmts[3*p+2]

	for i := range idxs {

		if lmts[3*i] < lmts[3*r] {

			idxs[l], idxs[i] = idxs[i], idxs[l]
			lmts[3*l], lmts[3*l+1], lmts[3*l+2], lmts[3*i], lmts[3*i+1], lmts[3*i+2] = lmts[3*i], lmts[3*i+1], lmts[3*i+2], lmts[3*l], lmts[3*l+1], lmts[3*l+2]

			l++

		}

	}

	idxs[l], idxs[r] = idxs[r], idxs[l]
	lmts[3*l], lmts[3*l+1], lmts[3*l+2], lmts[3*r], lmts[3*r+1], lmts[3*r+2] = lmts[3*r], lmts[3*r+1], lmts[3*r+2], lmts[3*l], lmts[3*l+1], lmts[3*l+2]

	sort(lmts[:3*l], idxs[:l])
	sort(lmts[3*l+3:], idxs[l+1:])

}
