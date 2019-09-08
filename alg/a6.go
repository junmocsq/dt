package alg

import (
	"errors"
	"fmt"
)

// 最大堆 除了根节点 a[parent[i]]>=a[i]
// 最小堆 除了根节点 a[parent[i]]<=a[i]
type HeapSort interface {
}

// 堆
type THeapSort struct {
	length int
	size   int
	array  []int
}

func getLeft(i int) int {
	return i<<1 + 1
}

func getRight(i int) int {
	return i<<1 + 2
}

func getParent(i int) int {
	return (i - 1) >> 1
}

// 下滤
func (this *THeapSort) MaxHeapify(i int) {
	largest := i
	for {
		l := getLeft(largest)
		r := getRight(largest)
		if l < this.size && this.array[largest] < this.array[l] {
			largest = l
		}
		if r < this.size && this.array[largest] < this.array[r] {
			largest = r
		}
		if largest != i {
			this.array[i], this.array[largest] = this.array[largest], this.array[i]
			i = largest
		} else {
			break
		}
	}
}

func (this *THeapSort) MinHeapify(i int) {
	smallest := i
	for {
		l := getLeft(i)
		r := getRight(i)
		if l < this.size && this.array[smallest] > this.array[l] {
			smallest = l
		}
		if r < this.size && this.array[smallest] > this.array[r] {
			smallest = r
		}
		if smallest != i {
			this.array[i], this.array[smallest] = this.array[smallest], this.array[i]
			i = smallest
		} else {
			break
		}
	}
}

func (this *THeapSort) Output() {
	fmt.Println("size:", this.size, "cap:", this.length)
	for i := 0; i < this.size; i++ {
		fmt.Printf("%d\t", this.array[i])
	}
	fmt.Printf("\n")
}

// 批量建堆
func (this *THeapSort) BuildMaxHeap(arr []int) {
	length := len(arr)
	this.array = arr
	this.length = 2 * length
	this.size = length
	for i := getParent(length - 1); i >= 0; i-- {
		this.MaxHeapify(i)
	}
}

func (this *THeapSort) BuildMinHeap(arr []int) {
	length := len(arr)
	this.array = arr
	this.length = 2 * length
	this.size = length
	for i := getParent(length - 1); i >= 0; i-- {
		this.MinHeapify(i)
	}
}

// 最大堆排序
func (this *THeapSort) MaxHeapSort() []int {
	var temp = &THeapSort{size: this.size, length: this.length, array: make([]int, this.size)}
	copy(temp.array, this.array)
	for i := temp.size - 1; i > 0; i-- {
		temp.array[i], temp.array[0] = temp.array[0], temp.array[i]
		temp.size--
		temp.MaxHeapify(0)
	}
	return temp.array
}

// 最小堆排序
func (this *THeapSort) MinHeapSort() []int {
	var temp = &THeapSort{size: this.size, length: this.length, array: make([]int, this.size)}
	copy(temp.array, this.array)
	for i := temp.size - 1; i > 0; i-- {
		temp.array[i], temp.array[0] = temp.array[0], temp.array[i]
		temp.size--
		temp.MinHeapify(0)
	}
	return temp.array
}

func NewHeap(length int) *THeapSort {
	return &THeapSort{size: 0, length: length, array: make([]int, length)}
}

func (this *THeapSort) check(index int) bool {
	if index < 0 || index >= this.size {
		return false
	} else {
		return true
	}
}

// 底层数组扩容
func (this *THeapSort) expand() {
	if this.size == this.length {
		var length int
		if this.length == 0 {
			length = 4
		} else if this.length > 1024 {
			length = 5 * this.length / 4
		} else {
			length = 2 * this.length
		}
		newHeap := &THeapSort{size: this.size, length: length, array: make([]int, length)}
		for i := 0; i < this.size; i++ {
			newHeap.array[i] = this.array[i]
		}
		*this = *newHeap
	}
}

// 底层数组缩容
func (this *THeapSort) shrink() {
	if this.size*4 <= this.length {
		var length int
		if this.length <= 4 {
			return
		} else {
			length = this.length / 4
			if length < 4 {
				length = 4
			}
		}
		newHeap := &THeapSort{size: this.size, length: length, array: make([]int, length)}
		for i := 0; i < this.size; i++ {
			newHeap.array[i] = this.array[i]
		}
		*this = *newHeap
	}
}

// 最大堆插入 插入最后采用上滤
func (this *THeapSort) MaxHeapInsert(val int) {
	this.expand()
	this.array[this.size] = val
	this.size++
	// 上滤
	curr := this.size - 1
	this.UpperHeapify(curr)
}

func (this *THeapSort) UpperHeapify(index int) {
	for this.size > 1 && index > 0 {
		parent := getParent(index)
		if this.array[parent] > this.array[index] {
			break
		} else {
			this.array[parent], this.array[index] = this.array[index], this.array[parent]
		}
		if parent == 0 {
			break
		}
		index = parent
	}
}

// 去除最大元素并返回
func (this *THeapSort) HeapExtractMax() (int, error) {
	if this.size == 0 {
		return 0, errors.New("优先队列为空")
	}
	max := this.array[0]
	this.array[0] = this.array[this.size-1]
	this.MaxHeapify(0)
	this.size--
	this.shrink()
	return max, nil
}

func (this *THeapSort) HeapIncreaseKey(index, incrNum int) bool {
	if !this.check(index) {
		return false
	}
	this.array[index] += incrNum
	this.UpperHeapify(index)
	return true
}

// 返回优先队列最大值
func (this *THeapSort) HeapMaxImum() (int, error) {
	if this.size == 0 {
		return 0, errors.New("优先队列为空")
	}
	return this.array[0], nil
}
