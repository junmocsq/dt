package alg

import (
	"fmt"
	"testing"
)

func TestTHeapSort_MaxHeapify(t *testing.T) {
	h := &THeapSort{length: 0, size: 0, array: []int{}}
	arr := make([]int, 20)
	for i := 0; i < 20; i++ {
		arr[i] = i + 1
	}
	arr = []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	h.BuildMaxHeap(arr)
	h.Output()
	fmt.Println(h.MaxHeapSort())
	h.BuildMinHeap(arr)
	h.Output()
	fmt.Println(h.MinHeapSort())
	h.array = []int{5, 13, 2, 25, 7, 17, 20, 8, 4}
	h.size = 9
	h.length = 20
	fmt.Println(h.MaxHeapSort())

}
