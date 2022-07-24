package main

import "github.com/duysmile/go-util/bst"

func main() {
	tree := &bst.BST{}

	tree.Add(3)
	tree.Add(6)
	tree.Add(2)
	tree.Add(4)
	tree.Add(5)
	tree.Add(1)
	tree.Display()

	tree.Remove(6)
	tree.Display()
}
