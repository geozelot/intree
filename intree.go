package intree

import (
	"math"
	"math/rand"
)


type GenericBounds interface {

  Limits() (float64, float64)

}


type Bounds struct {

  Upper, Lower float64

}

  func (bounds *Bounds) Limits() (float64, float64) {

    return bounds.Upper, bounds.Lower

  }


type INTree struct {

  nodeSize int

  idxs     []int
  lmts     []float64

}

  func (inT *INTree) buildIndex(bounds []GenericBounds, nodeSize int) {
    
    inT.nodeSize = nodeSize

    inT.idxs = make([]int, len(bounds))
    inT.lmts = make([]float64, 2*len(bounds))

    for i, v := range bounds {

      inT.idxs[i] = i
      u, l := v.Limits()

      inT.lmts[i*2] = l
      inT.lmts[i*2+1] = u

    }

    sort(inT.lmts, inT.idxs, inT.nodeSize)

  }

  func (inT *INTree) Includes(val float64) []int {

    stack  := []int{0, len(inT.idxs) - 1}
    result := []int{}

    var u, l float64

    for len(stack) > 0 {

      right:= stack[len(stack)-1]
      stack = stack[:len(stack)-1]
      left := stack[len(stack)-1]
      stack = stack[:len(stack)-1]

      if right-left <= inT.nodeSize {

        for i := left; i <= right; i++ {

          l = inT.lmts[2*i]
          u = inT.lmts[2*i+1]

          if u >= val && l <= val {

            result = append(result, inT.idxs[i])

          }

        }

        continue
      
      }

      m := int(math.Floor(float64(left + right) / 2.0))

      l = inT.lmts[2*m]
      u = inT.lmts[2*m+1]

      if u >= val && l <= val {

        result = append(result, inT.idxs[m])

      }

      if (val <= u) {

        stack = append(stack, left)
        stack = append(stack, m - 1)

      } else {

        stack = append(stack, m + 1)
        stack = append(stack, right)

      }

    }

    return result
    
  }


func sort(lmts []float64, idxs []int, nodeSize int) {

  if len(lmts) < 2 { return }
    
  left, right := 1, len(lmts) - 1

  if (right - left) <= nodeSize { return }
    
  pivot := rand.Int() % len(lmts)

  if pivot % 2 == 0 { pivot++ }
    
  lmts[pivot], lmts[pivot-1], lmts[right], lmts[right-1] = lmts[right], lmts[right-1], lmts[pivot], lmts[pivot-1]
  idxs[(pivot-1)/2], idxs[(right-1)/2] = idxs[(right-1)/2], idxs[(pivot-1)/2]

  for i := 1; i < len(lmts); i += 2  {

      if lmts[i] < lmts[right] {

        lmts[left], lmts[left-1], lmts[i], lmts[i-1] = lmts[i], lmts[i-1], lmts[left], lmts[left-1]
        idxs[(left-1)/2], idxs[(i-1)/2] = idxs[(i-1)/2], idxs[(left-1)/2]
        
        left += 2
      
      }
  
  }

  lmts[left], lmts[left-1], lmts[right], lmts[right-1] = lmts[right], lmts[right-1], lmts[left], lmts[left-1]
  idxs[(left-1)/2], idxs[(right-1)/2] = idxs[(right-1)/2], idxs[(left-1)/2]
    
  sort(lmts[:left-1], idxs[:(left-1)/2], nodeSize)
  sort(lmts[left+1:], idxs[(((left-1)/2)+1):], nodeSize)
    
  return

}


func NewINTree(bounds []GenericBounds, nodeSize ...int) *INTree {

	nSize := nodeSize[0]

	if nSize == 0 {
		nSize = 64
	}

  b := INTree{}
  b.buildIndex(bounds, nSize)

  return &b

}