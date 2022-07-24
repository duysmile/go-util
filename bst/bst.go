package bst

import (
	"fmt"
	"math"
)

type Side int

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	value  int
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) SetLeft(left *Node) {
	if left != nil {
		left.SetParent(n)
	}
	n.left = left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) SetRight(right *Node) {
	if right != nil {
		right.SetParent(n)
	}
	n.right = right
}

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) SetParent(parent *Node) {
	n.parent = parent
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) SetValue(value int) {
	n.value = value
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

type BST struct {
	root *Node
}

func (t *BST) Root() *Node {
	return t.root
}

func (t *BST) SetRoot(node *Node) {
	t.root = node
}

func (t *BST) Add(value int) *Node {
	node := &Node{
		value: value,
	}

	if t.root == nil {
		t.root = node
		return node
	}

	currentNode := t.root
	for {
		if currentNode.value < node.value {
			if currentNode.right == nil {
				node.parent = currentNode
				currentNode.right = node
				break
			} else {
				currentNode = currentNode.right
			}
		} else if currentNode.value > node.value {
			if currentNode.left == nil {
				node.parent = currentNode
				currentNode.left = node
				break
			} else {
				currentNode = currentNode.left
			}
		}
	}

	return node
}

func (t *BST) Find(value int) *Node {
	node := t.root
	for {
		if node == nil || node.value == value {
			return node
		}

		if node.value > value {
			node = node.left
		} else {
			node = node.right
		}
	}
}

func (t *BST) Remove(value int) bool {
	removeNode := t.Find(value)
	if removeNode == nil {
		return false
	}

	nodeToReplace := t.getReplaceNode(removeNode)
	nodeToReplace.parent = nil

	if t.root == removeNode {
		t.root = nodeToReplace
		t.root.parent = nil
	} else {
		parent := removeNode.parent
		if parent.left == removeNode {
			parent.left = nodeToReplace
		} else {
			parent.right = nodeToReplace
		}
	}

	return true
}

// getReplaceNode get and adjust node to replace a removed node
func (t *BST) getReplaceNode(node *Node) *Node {
	if node.right != nil {
		leftMostNode := t.getLeftMostNode(node.right)
		leftMostNode.left = node.left
		return node.right
	}

	return node.left
}

func (t *BST) getLeftMostNode(node *Node) *Node {
	for {
		if node.left == nil {
			return node
		}

		node = node.left
	}
}

func (t *BST) GetHeight(node *Node) int {
	if node == nil {
		return -1
	}

	height := math.Max(
		float64(t.GetHeight(node.left)+1),
		float64(t.GetHeight(node.right)+1),
	)
	return int(height)
}

func (t *BST) GetLevel(node *Node) int {
	if node == nil {
		return -1
	}

	level := 0
	for {
		parent := node.parent
		if parent == nil {
			return level
		}
		level++
		node = parent
	}
}

func (t *BST) Display() {
	listNode := make([]*Node, 0)

	if t.root == nil {
		return
	}

	listNode = append(listNode, t.root)
	level := 0
	for {
		if len(listNode) == 0 {
			fmt.Println("")
			return
		}

		currentNode := listNode[0]
		listNode = listNode[1:]

		if level < t.GetLevel(currentNode) {
			fmt.Println("")
			level++
		}

		fmt.Printf("%v[height: %v] ", currentNode.value, t.GetHeight(currentNode))
		if currentNode.left != nil {
			listNode = append(listNode, currentNode.left)
		}
		if currentNode.right != nil {
			listNode = append(listNode, currentNode.right)
		}
	}
}
