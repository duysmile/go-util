package main

import (
	"fmt"
	"github.com/duysmile/go-util/avl"
	"github.com/duysmile/go-util/bst"
)

func main() {
	// Left rotation
	n1 := bst.NewNode(1)
	n2 := bst.NewNode(2)
	n4 := bst.NewNode(4)
	n3 := bst.NewNode(3)
	n5 := bst.NewNode(5)

	n1.SetRight(n2)
	n2.SetRight(n4)
	n4.SetRight(n5)
	n4.SetLeft(n3)

	fmt.Println(avl.LeftRotation(n2) == n4)

	// LeftRightRotation
	no1 := bst.NewNode(1)
	no2 := bst.NewNode(2)
	no3 := bst.NewNode(3)

	no3.SetLeft(no1)
	no1.SetRight(no2)

	fmt.Println(avl.LeftRightRotation(no3) == no2)

	// AVL Tree
	tree := &avl.AVL{}
	tree.Add(3)
	tree.Add(4)
	tree.Add(0)
	tree.Add(2)
	tree.Add(1)
	tree.Add(-1)

	tree.Display()
}
