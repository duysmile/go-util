package main

import (
	"fmt"
	"github.com/duysmile/go-util/consistenthash"
)

func main() {
	hash := consistenthash.NewHash(2, nil)

	hash.AddMulti("1", "2", "3")

	fmt.Println(hash.Get("abc"))
	fmt.Println(hash.Get("abd"))
	fmt.Println(hash.Get("abe"))
	fmt.Println(hash.Get("abf"))
	fmt.Println(hash.Get("abg"))
	fmt.Println(hash.Get("abh"))

	fmt.Println("remove node 1")
	hash.Remove("1")

	fmt.Println(hash.Get("abc"))
	fmt.Println(hash.Get("abd"))
	fmt.Println(hash.Get("abe"))
	fmt.Println(hash.Get("abf"))
	fmt.Println(hash.Get("abg"))
	fmt.Println(hash.Get("abh"))
}
