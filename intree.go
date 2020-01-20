package intree

import (

	"fmt"

)


type Bounds struct {

   Lower, Upper int

}

type tNode struct {

	indices				[]int
  bounds				*Bounds
  max						int
  left, right		*tNode

}

type INTree struct {
	
	root		*tNode
	index		int

}


func (inTree *INTree) newNode(bounds *Bounds) *tNode {

	node := &tNode {

		indices:	[]int{inTree.index},
		bounds:	&Bounds {
			Lower:	bounds.Lower,
			Upper:	bounds.Upper,
		},
		max:		bounds.Upper,
		left:		nil,
		right:	nil,

	}

	inTree.index++

	return node

}

func (inTree *INTree) insertNode(currentNode *tNode, bounds *Bounds) *tNode {

	if (currentNode == nil) {
	  return inTree.newNode(bounds)
	}

	if (currentNode.bounds.Lower == bounds.Lower && currentNode.bounds.Upper == bounds.Upper) {

		currentNode.indices = append(currentNode.indices, inTree.index)

		inTree.index++

		return currentNode

	}

	currentNode.indices = append(currentNode.indices, inTree.index)

	low := currentNode.bounds.Lower

	if (bounds.Lower < low) {

	  currentNode.left = inTree.insertNode(currentNode.left, bounds)

	} else {

	  currentNode.right = inTree.insertNode(currentNode.right, bounds)

	}

	if (currentNode.max < bounds.Upper) {

	  currentNode.max = bounds.Upper

	}

	return currentNode

}

func (inTree *INTree) includesValue(node *tNode, value int) []int {

	if (node == nil) {
		var indices []int
	  return indices
	}

	if (node.bounds.Lower <= value && node.bounds.Upper >= value) {
/*
		indices := node.indices
		indices = append(indices, inTree.getChildNodeIndices(node.left)...)
		indices = append(indices, inTree.getChildNodeIndices(node.right)...)
*/
	  return node.indices

	}

	if (node.left != nil && node.left.max >= value) {
	  return inTree.includesValue(node.left, value)
	}

	return inTree.includesValue(node.right, value)

}

func (inTree *INTree) inOrder(node *tNode) {

	if (node == nil) {
  	return
  }
  
  inTree.inOrder(node.left)

  fmt.Println(node.bounds, node.max, node.indices)

  inTree.inOrder(node.right)

}
/*
func (inTree *INTree) getChildNodeIndices(node *tNode) []int {

	if (node == nil) {
		return make([]int, 0)
	}
	
	indices := node.indices
	indices = append(indices, inTree.getChildNodeIndices(node.left)...)
	indices = append(indices, inTree.getChildNodeIndices(node.right)...)

	return indices

}
*/
func NewINTree() *INTree {

	tree := INTree {

		root: 	nil,
		index: 	0,

	}

	return &tree

}

func (inTree *INTree) InOrder() {

	inTree.inOrder(inTree.root)

}

func (inTree *INTree) Includes(value int) []int {

	return inTree.includesValue(inTree.root, value)

}

func (inTree *INTree) Insert(bounds *Bounds) {

	inTree.root = inTree.insertNode(inTree.root, bounds)

}