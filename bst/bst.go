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

type BST struct {
	root *Node
}

func (t *BST) Add(value int) {
	node := &Node{
		value: value,
	}

	if t.root == nil {
		t.root = node
		return
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

func (t *BST) getHeight(node *Node) int {
	if node == nil {
		return -1
	}

	height := math.Max(
		float64(t.getHeight(node.left)+1),
		float64(t.getHeight(node.right)+1),
	)
	return int(height)
}

func (t *BST) getLevel(node *Node) int {
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

		if level < t.getLevel(currentNode) {
			fmt.Println("")
			level++
		}

		fmt.Printf("%v[height: %v] ", currentNode.value, t.getHeight(currentNode))
		if currentNode.left != nil {
			listNode = append(listNode, currentNode.left)
		}
		if currentNode.right != nil {
			listNode = append(listNode, currentNode.right)
		}
	}
}
