package alg

import "fmt"

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

func (this *THeapSort) MaxHeapInsert() {

}

func (this *THeapSort) HeapExtractMax() {

}

func (this *THeapSort) HeapIncreaseKey() {

}

func (this *THeapSort) HeapMaxImum() {

}
