package part4

import (
	"fmt"
	"testing"
)

func TestBinTree(t *testing.T) {
	tree := NewTree()
	arr := []int{4, 2, 6, 1, 3, 5, 7}
	for _, v := range arr {
		tree.Insert(v)
	}
	tree.PrintTree()
	fmt.Println(tree.Delete(4))
	fmt.Println(tree.Delete(2))
	fmt.Println(tree.Delete(6))
	tree.PrintTree()
}
