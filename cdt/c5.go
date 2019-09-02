package cdt

import (
	"errors"
	"fmt"
)

// 线性表
type LinearList interface {
	Empty() bool
	Size() int
	Get(int) (int, error)
	IndexOf(int) int
	Erase(index int) (int, bool)
	Insert(index int, val int) bool
	Output()
}

const MIN_ARRAY_LEN = 4

type ArrLinearList struct {
	array  []int
	length int
	cap    int
}

func NewArr(length int) *ArrLinearList {
	temp := make([]int,length)
	return &ArrLinearList{array: temp, length: 0, cap: length}
}

// 复制
func (this *ArrLinearList) copy(arr *ArrLinearList) bool {
	if this.length > arr.cap {
		return false
	}
	for i := 0; i < this.length; i++ {
		arr.array[i] = this.array[i]
	}
	arr.length = this.length
	return true
}

// 扩容
func (this *ArrLinearList) expand() {
	if this.length == this.cap {
		length := 2 * this.length
		newArrLinearList := NewArr(length)
		if this.copy(newArrLinearList) {
			*this = *newArrLinearList
		}
	}
}

// 缩容
func (this *ArrLinearList) shrink() {
	if this.length*4 <= this.cap && this.cap > MIN_ARRAY_LEN {
		length := this.cap / 2
		if length < MIN_ARRAY_LEN {
			length = MIN_ARRAY_LEN
		}
		newArrLinearList := NewArr(length)
		this.copy(newArrLinearList)
		if this.copy(newArrLinearList) {
			*this = *newArrLinearList
		}
	}
}

func (this *ArrLinearList) Empty() bool {
	if this.length == 0 {
		return true
	} else {
		return false
	}
}

func (this *ArrLinearList) Size() int {
	return this.length
}

func (this *ArrLinearList) Get(index int) (int, error) {
	if this.length <= index || index < 0 {
		return 0, errors.New("数组越界")
	}
	return this.array[index], nil
}

// 查找索引 不存在返回-1
func (this *ArrLinearList) IndexOf(val int) int {
	res := -1
	for i := 0; i < this.length; i++ {
		if this.array[i] == val {
			res = i
			break
		}
	}
	return res
}

// 删除
func (this *ArrLinearList) Erase(index int) (int, error) {
	if this.length <= index || index < 0 {
		return 0, errors.New("数组越界")
	}
	temp := this.array[index]
	for i := index; i < this.length-1; i++ {
		this.array[i] = this.array[i+1]
	}
	this.length--
	this.shrink() //缩容
	return temp, nil
}

// 插入
func (this *ArrLinearList) Insert(index int, val int) bool {
	if this.length < index || index < 0 {
		return false
	}
	this.expand() // 扩容
	this.length++
	for i := this.length - 1; i > index; i-- {
		this.array[i] = this.array[i-1]
	}
	this.array[index] = val
	return true
}

func (this *ArrLinearList) Output() {
	fmt.Println("length:", this.length, " cap:", this.cap)
	for i := 0; i < this.length; i++ {
		fmt.Printf("%d \t", this.array[i])
	}
	fmt.Printf("\n")
}
