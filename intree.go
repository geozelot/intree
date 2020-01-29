// INTree provides a static, flat (augmented) INterval Tree for reverse range searches.
package intree

import (
	"math"
	"math/rand"
)


// Bounds is the main interface expected by NewINTree();
type Bounds interface {
	Limits() (L, U float64)
}

// SimpleBounds is a simple Struct implicitly implementing the Bounds interface.
type SimpleBounds struct {
	Lower, Upper float64
}

// Limits implicitly implements Bounds interface for Simplebounds
func (sb *SimpleBounds) Limits() (float64, float64) {

	return sb.Lower, sb.Upper

}

// INTree is the main package object.
// INTree.idxs holds the array indices of the passed Bounds[] Slice at construction with NewINTree(),
// INTree.lmts holds the actual limits, and the augmented maximum of all children per node.
type INTree struct {
	idxs []int
	lmts []float64
}

// buidTree is the internal tree construction function.
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

// Including is the main public entry point for bounds searches.
// It will traverse the tree and return overlaps with the given value.
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

// NewINTree is the main initialization function.
// It creates the tree from all passed in Bounds objects by calling buildTree.
func NewINTree(bnds []Bounds) *INTree {

	inT := INTree{}
	inT.buildTree(bnds)

	return &inT

}

// augment is a internal utility function to add maximum value of all child nodes to the current node.
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

	r := int(math.Floor(float64(len(idxs)) / 2.0))

	lmts[3*r+2] = max

	augment(lmts[:3*r], idxs[:r])
	augment(lmts[3*r+3:], idxs[r+1:])

}

// sort is a internal utility function to sort tree by lowest limits, using Random Pivot QuickSearch
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
