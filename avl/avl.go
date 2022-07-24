package avl

import "github.com/duysmile/go-util/bst"

/* A tree is balanced if:
- #1 The left subtree height and the right subtree height differ by at most 1.
- #2 Visit every node making sure rule #1 is satisfied.
Note: Height of a node is the distance (edge count) from the farthest child to itself.
References: https://adrianmejia.com/self-balanced-binary-search-trees-with-avl-tree-data-structure-for-beginners/
*/

func resetParent(parent, oldChild, newChild *bst.Node) {
	if parent != nil {
		if parent.Left() == oldChild {
			parent.SetLeft(newChild)
		} else {
			parent.SetRight(newChild)
		}
	} else {
		newChild.SetParent(nil)
	}

	oldChild.SetParent(newChild)
}

func LeftRotation(node *bst.Node) *bst.Node {
	if node == nil {
		return nil
	}

	right := node.Right()
	if right == nil {
		return nil
	}
	resetParent(node.Parent(), node, right)
	right.SetLeft(node)
	node.SetRight(nil)
	return right
}

func RightRotation(node *bst.Node) *bst.Node {
	if node == nil {
		return nil
	}

	left := node.Left()
	if left == nil {
		return nil
	}
	resetParent(node.Parent(), node, left)
	left.SetRight(node)
	node.SetLeft(nil)
	return left
}

// LeftRightRotation first LeftRotation on left of node and then RightRotation on node
func LeftRightRotation(node *bst.Node) *bst.Node {
	if node == nil {
		return nil
	}

	left := node.Left()
	if left == nil {
		return nil
	}

	lNode := LeftRotation(left)
	if lNode == nil {
		return nil
	}

	return RightRotation(node)
}

// RightLeftRotation first RightRotation on right of node and then LeftRotation on node
func RightLeftRotation(node *bst.Node) *bst.Node {
	if node == nil {
		return nil
	}

	right := node.Right()
	if right == nil {
		return nil
	}

	rNode := RightRotation(right)
	if rNode == nil {
		return nil
	}

	return LeftRotation(node)
}

type AVL struct {
	bst.BST
}

// balanceFactor calculates subtraction of left-subtree's height and right-subtree's height
func (t *AVL) balanceFactor(node *bst.Node) int {
	return t.GetHeight(node.Left()) - t.GetHeight(node.Right())
}

func (t *AVL) balance(node *bst.Node) *bst.Node {
	nodeBalanceFactor := t.balanceFactor(node)

	if nodeBalanceFactor > 1 {
		leftBalanceFactor := t.balanceFactor(node.Left())
		if leftBalanceFactor > 0 {
			return RightRotation(node)
		} else if leftBalanceFactor < 0 {
			return LeftRightRotation(node)
		}
	} else if nodeBalanceFactor < -1 {
		rightBalanceFactor := t.balanceFactor(node.Right())
		if rightBalanceFactor < 0 {
			return LeftRotation(node)
		} else if rightBalanceFactor > 0 {
			return RightLeftRotation(node)
		}
	}

	return node
}

func (t *AVL) Add(value int) *bst.Node {
	node := t.BST.Add(value)
	t.SetRoot(t.balanceUpstream(node))
	return node
}

func (t *AVL) Remove(value int) bool {
	removeNode := t.Find(value)
	if removeNode != nil {
		t.BST.Remove(value)
		t.SetRoot(t.balanceUpstream(removeNode.Parent()))
		return true
	}

	return false
}

func (t *AVL) balanceUpstream(node *bst.Node) *bst.Node {
	currentNode := node
	var newParent *bst.Node
	for {
		if currentNode == nil {
			return newParent
		}
		newParent = t.balance(currentNode)
		currentNode = currentNode.Parent()
	}
}
