package intree

import (

  "math"
  "math/rand"

)



/*
 *  Main interface expected by NewINTree();
 *
 *    @member Limits(): returns (multiple-value) lower and upper limit of implementing object
 */

type Bounds interface {

  Limits() (L, U float64)

}


/*
 *  Simple struct implicitly implementing Bounds interface;
 *
 *    @property Upper:  holds upper limit of bounds
 *    @property Lower:  holds lower limit of bounds
 *
 *    @method Limits(): returning Upper and Lower properties
 */

type SimpleBounds struct {

  Lower, Upper float64

}


  /*
   *  Implementing Bounds interface method
   *
   *    @return:  returns lower and upper limit of bounds
   */

  func (sb *SimpleBounds) Limits() (float64, float64) {

    return sb.Lower, sb.Upper

  }


/*
 *  Main INTree object;
 *
 *    @property idxs:     indices referencing index positions of the []Bounds array passed in to construct the tree
 *    @property lmts:     { lower limit (lmts[3*i]); upper limit (lmts[3*i+1]); maxmimum value of left/right child nodes (lmts[3*i+2]) }
 *
 *    @method buildTree:  internal tree construction function; called by NewINTree(); calls utility functions sort() and augment() to build node dependencies
 *    @method Including:  main public entry point: finds all bounds that include the given value
 */ 

type INTree struct {

  idxs     []int
  lmts     []float64

}


  /*
   *  Internal tree construction function
   *    
   *    @param bnds: Slice of objects implementing Bounds[] interface
   */

  func (inT *INTree) buildTree(bnds []Bounds) {
    
    inT.idxs = make([]int, len(bnds))
    inT.lmts = make([]float64, 3*len(bnds))

    for i, v := range bnds {

      inT.idxs[i] = i
      l, u := v.Limits()

      inT.lmts[3*i]   = l
      inT.lmts[3*i+1] = u
      inT.lmts[3*i+2] = 0

    }

    sort(inT.lmts, inT.idxs)
    augment(inT.lmts, inT.idxs)

  }


  /*
   *  Main public entry point for bounds searches
   *
   *    @param val: value to search containing bounds for
   *
   *    @return:    Slice holding all found array indices referencing the input Bounds[]
   */

  func (inT *INTree) Including(val float64) []int {

    stk := []int{0, len(inT.idxs) - 1}
    res := []int{}

    for len(stk) > 0 {

      rb  := stk[len(stk)-1]
      stk  = stk[:len(stk)-1]
      lb  := stk[len(stk)-1]
      stk  = stk[:len(stk)-1]

      if lb == rb + 1 { continue }

      cn := int(math.Ceil(float64(lb + rb) / 2.0))
      nm := inT.lmts[3*cn+2]

      if val <= nm {

        stk = append(stk, lb)
        stk = append(stk, cn - 1)

      }

      l := inT.lmts[3*cn]

      if l <= val {

        stk = append(stk, cn + 1)
        stk = append(stk, rb)

        u := inT.lmts[3*cn+1]
        
        if val <= u { res = append(res, inT.idxs[cn]) }

      }

    }

    return res
  
  }



/*
 *  Main initialization function; creates the tree from passed in Bounds objects by calling method buildTree
 *
 *    @params bnds: Slice of objects implementing the Bounds interface
 *
 *    @return:      INTree object
 */

func NewINTree(bnds []Bounds) *INTree {

  inT := INTree{}
  inT.buildTree(bnds)

  return &inT

}



/*
 *  Utility function to augment tree by adding maximum value of all child nodes
 *
 *    @param lmts: Slice partition of (lower, upper, max) values defining the tree nodes
 *    @param idxs: Slice partition of (index) values referencing input collection of Bounds objects
 */

func augment(lmts []float64, idxs []int) {

  if len(idxs) < 1 { return }

  max := 0.0

  for idx, _ := range idxs {

    if lmts[3*idx+1] > max { max = lmts[3*idx+1] }

  }

  r := int(math.Floor(float64(len(idxs)) / 2.0))

  lmts[3*r+2] = max

  augment(lmts[:3*r], idxs[:r])
  augment(lmts[3*r+3:], idxs[r+1:])

}


/*
 *  Utility function to sort tree by lowest limits
 *
 *    @param lmts: Slice partition of { lower, upper, max } values defining the tree nodes
 *    @param idxs: Slice partition of { index } values referencing input collection of Bounds objects
 */

func sort(lmts []float64, idxs []int) {

  if len(idxs) < 2 { return }
    
  l, r := 0, len(idxs) - 1

  p := rand.Int() % len(idxs)

  idxs[p], idxs[r] = idxs[r], idxs[p]
  lmts[3*p], lmts[3*p+1], lmts[3*p+2], lmts[3*r], lmts[3*r+1], lmts[3*r+2] = lmts[3*r], lmts[3*r+1], lmts[3*r+2], lmts[3*p], lmts[3*p+1], lmts[3*p+2]

  for i := range idxs  {

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