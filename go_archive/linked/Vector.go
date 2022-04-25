package linked

import (
	"dt/derror"
	"fmt"
	"math/rand"
)

const (
	ASC  = 1
	DESC = 2
)

type vector struct {
	arr      []int
	size     int
	capacity int
}

func VCreate(length int) *vector {
	return &vector{arr: make([]int, length), size: 0, capacity: length}
}

func (vec *vector) VCopyFrom(newVector *vector, lo int, hi int) bool {
	if vec.size < hi {
		return false
	}
	size := hi - lo
	if newVector.capacity < size {
		return false
	}
	var index int = 0
	for lo < hi {
		newVector.arr[index] = vec.arr[lo]
		index++
		lo++
	}
	return true
}

func (vec *vector) VSize() int {
	return vec.size
}

func (vec *vector) VEmpty() bool {
	return vec.size == 0
}

// 判断向量是否已排序 返回逆序对
func (vec *vector) VDisordered(sort int) int {
	var i int
	var count = 0
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

// 无序向量查找 返回查到的rank最大的值，如果没查到 hi为比lo小1的值
func (vec *vector) VFind(val int, lo int, hi int) (int, error) {
	for lo < hi && hi <= vec.size {
		hi--
		if val == vec.arr[hi] {
			return hi, nil
		}
	}
	return hi, derror.NewErr(4003)

}

func (vec *vector) VAdd(rank int, val int) bool {
	if rank > vec.size {
		rank = vec.size
	}
	vec.VExpand()
	var i int
	for i = vec.size; i > rank; i-- {
		vec.arr[i] = vec.arr[i-1]
	}
	vec.arr[rank] = val
	vec.size++
	return true

}

// 有序向量查找
func (vec *vector) VSearch() {

}

//
func (vec *vector) VRemove(lo int, hi int) bool {
	if lo < 0 || hi > vec.size || lo >= hi {
		return false
	}
	move := hi - lo
	var i int = lo
	for i < vec.size-move {
		vec.arr[i] = vec.arr[i+move]
		i++
	}
	vec.size -= move
	vec.VShrink() // 缩容
	return true
}

func (vec *vector) VDelete(rank int) bool {
	return vec.VRemove(rank, rank+1)
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
		*vec = *newVector

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
			// 指针对应值修改
			*vec = *newVector
			return true
		}
	}
	return false
}

// 随机置乱算法
func (vec *vector) VShuffle() {
	var i int
	for i = vec.size - 1; i > 1; i-- {
		vec.VBubble(i, rand.Intn(i-1))
	}
}

// 扫描交换
func (vec *vector) VBubble(lo int, hi int) bool {
	if vec.size <= hi || lo >= hi {
		return false
	}
	ele := vec.arr[lo]
	vec.arr[lo] = vec.arr[hi]
	vec.arr[hi] = ele
	return true
}

// 冒泡排序 相邻两个元素比较
func (vec *vector) VBubbleSort(lo int, hi int) bool {
	if vec.size < hi {
		return false
	}
	var i, j int
	for i = lo; i < hi-1; i++ {
		for j = lo + 1; j < hi-i; j++ {
			if vec.arr[j-1] < vec.arr[j] {
				vec.VBubble(j-1, j)
			}
		}
	}
	return true
}

// 数据去重
func (vec *vector) VDeduplicate() int {
	oldSize := vec.size
	rank := 1
	for rank < vec.size {
		i, err := vec.VFind(vec.arr[rank-1], rank, vec.size)
		if err == nil {
			// 如果有重复数据 则需要再查找一遍确定无此重复数据 防止漏删 超过2个以上的重复数据
			vec.VDelete(i)
		} else {
			rank++
		}
	}
	return oldSize - vec.size
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

func (vec *vector) VPrint() {
	fmt.Println("[")
	for i := 0; i < vec.size; i++ {
		fmt.Println("	", i, "=>", vec.arr[i])
	}
	fmt.Println("]")
	fmt.Println("size:", vec.size, "capacity:", vec.capacity)
}
