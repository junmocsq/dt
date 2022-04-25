package alg

import (
	"fmt"
	"testing"
)

func TestTQuickSort(t *testing.T) {
	a := []int{3, 4, 2, 14, 42, 12, 11, 77, 4, 11, 7, 9}
	m := Alg7{}
	m.QuickSort(a, 0, len(a))
	fmt.Println(m)
}
