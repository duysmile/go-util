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
