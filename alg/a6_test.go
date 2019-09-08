package alg

import (
	"fmt"
	"testing"
	"time"
)

func TestTHeapSort_MaxHeapify(t *testing.T) {
	h := &THeapSort{length: 0, size: 0, array: []int{}}
	ti := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		temp := i + 1
		h.MaxHeapInsert(temp)
	}
	ti2 := time.Now().UnixNano()
	fmt.Println((ti2 - ti) / 1000000.0)

	ti3 := time.Now().UnixNano()
	fmt.Println((ti3 - ti2) / 1000000.0)
	h.Output()
	h.HeapIncreaseKey(0, 2)
	h.Output()
}
