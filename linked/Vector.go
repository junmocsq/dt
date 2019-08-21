package linked

import (
	"dt/derror"
	"math/rand"
)

const (
	ASC  = 1
	DESC = 2
)

type vector struct {
	arr      []int
	size     uint32
	capacity uint32
}

func VCreate(length uint32) *vector {
	return &vector{arr: make([]int, length), size: 0, capacity: length}
}

func (vec *vector) VCopyFrom(newVector *vector, lo uint32, hi uint32) bool {
	if vec.size <= hi {
		return false
	}
	size := hi - lo
	if newVector.capacity < size {
		return false
	}
	var index uint32 = 0
	for lo < hi {
		newVector.arr[index] = vec.arr[lo]
		index++
		lo++
	}
	return true
}

func (vec *vector) VSize() uint32 {
	return vec.size
}

func (vec *vector) VEmpty() bool {
	return vec.size == 0
}

// 判断向量是否已排序 返回逆序对
func (vec *vector) VDisordered(sort uint32) uint32 {
	var i uint32
	var count uint32 = 0
	for i = 1; i < vec.size; i++ {
		if sort == ASC {
			if vec.arr[i-1] > vec.arr[i] {
				count++
			}
		} else {
			if vec.arr[i-1] < vec.arr[i] {
				count++
			}
		}
	}
	return count
}

// 无序向量查找
func (vec *vector) VFind() {

}

// 有序向量查找
func (vec *vector) VSearch() {

}

//
func (vec *vector) VRemove() {

}

// 空间不足时扩容
func (vec *vector) VExpand() error {
	if vec.size < vec.capacity {
		return nil
	}
	newVector := VCreate(vec.capacity * 2)
	newVector.size = vec.size
	res := vec.VCopyFrom(newVector, 0, vec.size)
	if res {
		vec = newVector
		return nil
	} else {
		return derror.NewErr(4002)
	}
}

// 装填因子太小时压缩
func (vec *vector) VShrink() bool {
	if vec.size*4 < vec.capacity {
		newVector := VCreate(vec.capacity / 2)
		newVector.size = vec.size
		res := vec.VCopyFrom(newVector, 0, vec.size)
		if res {
			vec = newVector
			return true
		}
	}
	return false
}

// 随机置乱算法
func (vec *vector) VShuffle() {
	var i uint32
	for i = vec.size; i > 1; i-- {
		vec.VBubble(i, rand.Uint32()%(i-1))
	}
}

// 扫描交换
func (vec *vector) VBubble(lo uint32, hi uint32) bool {
	if vec.size <= hi || lo >= hi {
		return false
	}
	ele := vec.arr[lo]
	vec.arr[lo] = vec.arr[hi]
	vec.arr[hi] = ele
	return true
}

func (vec *vector) VBubbleSort(lo uint32, hi uint32) {

}

func (vec *vector) VMax() {

}

func (vec *vector) VSelectionSort() {

}

func (vec *vector) VMerge() {

}

func (vec *vector) VMergeSort() {

}

func (vec *vector) VPartition() {

}

func (vec *vector) VQuickSort() {

}

func (vec *vector) VHeapSort() {

}
