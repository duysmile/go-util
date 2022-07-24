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
	n3 := bst.NewNode(3)
	n4 := bst.NewNode(4)

	n1.SetRight(n2)
	n2.SetRight(n3)
	n3.SetRight(n4)

	fmt.Println(avl.LeftRotation(n2) == n3)

	// Right rotation
	n8 := bst.NewNode(8)
	n7 := bst.NewNode(7)
	n6 := bst.NewNode(6)
	n5 := bst.NewNode(5)

	n8.SetLeft(n7)
	n7.SetLeft(n6)
	n6.SetLeft(n5)

	fmt.Println(avl.RightRotation(n7) == n6)

	// LeftRightRotation
	no1 := bst.NewNode(1)
	no2 := bst.NewNode(2)
	no3 := bst.NewNode(3)

	no3.SetLeft(no1)
	no1.SetRight(no2)

	fmt.Println(avl.LeftRightRotation(no3) == no2)

	// RightLeftRotation
	no4 := bst.NewNode(4)
	no5 := bst.NewNode(5)
	no6 := bst.NewNode(6)

	no4.SetRight(no6)
	no6.SetLeft(no5)

	fmt.Println(avl.RightLeftRotation(no4) == no5)

	// AVL Tree
	tree := &avl.AVL{}
	tree.Add(3)
	tree.Add(1)
	tree.Add(2)

	tree.Display()
}
