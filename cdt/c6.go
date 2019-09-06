package cdt

import (
	"errors"
	"fmt"
)

type ChainNode struct {
	ele  int
	next *ChainNode
}

type Chain struct {
	head *ChainNode
	tail *ChainNode
	size int
}

/**
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
*/
func (this *Chain) Empty() bool {
	if this.size == 0 {
		return true
	} else {
		return false
	}
}

func (this *Chain) Size() int {
	return this.size
}

func (this *Chain) Get(index int) (int, error) {
	if this.size <= index || index < 0 {
		return 0, errors.New("访问的元素越界")
	}
	temp := this.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.ele, nil
}

func (this *Chain) IndexOf(val int) int {
	temp := this.head
	for i := 0; i < this.size; i++ {
		if temp.ele == val {
			return i
		}
		if temp.next != nil {
			temp = temp.next
		} else {
			break
		}
	}
	return -1
}

func (this *Chain) Erase(index int) (int, error) {
	if this.size <= index || index < 0 {
		return 0, errors.New("访问的元素越界")
	}
	var val int
	if index == 0 {
		val = this.head.ele
		this.head = this.head.next
		if this.head == nil {
			this.tail = nil
		}
	} else {
		temp := this.head
		for i := 0; i < index-1; i++ {
			temp = temp.next
		}
		val = temp.next.ele
		temp.next = temp.next.next
		if temp.next == nil {
			this.tail = temp
		}
	}
	this.size--
	return val, nil

}

func (this *Chain) Insert(index int, val int) bool {
	if index > this.size || index < 0 {
		return false
	}
	node := ChainNode{ele: val, next: nil}
	if index == 0 {
		node.next = this.head
		this.head = &node
	} else {
		temp := this.head
		for i := 0; i < index-1; i++ {
			temp = temp.next
		}
		node.next = temp.next
		temp.next = &node
	}
	if node.next == nil {
		this.tail = &node
	}
	this.size++
	return true
}

func (this *Chain) PushBack(val int) {
	node := ChainNode{ele: val, next: nil}
	if this.head == nil {
		this.head = &node
		this.tail = &node
	} else {
		this.tail.next = &node
		this.tail = &node
	}
	this.size++
}

func (this *Chain) Output() {
	fmt.Println("Chain size:", this.size)
	temp := this.head
	for i := 0; i < this.size; i++ {
		fmt.Printf("%d\t", temp.ele)
		if temp.next != nil {
			temp = temp.next
		}
	}
	fmt.Printf("\n")
}

func (this *Chain) Reverse() {
	if this.size <= 2 {
		if this.size == 2 {
			this.tail.next = this.head
			this.head.next = nil
		}
	} else {
		pre := this.head
		mid := pre.next
		pre.next = nil
		for {
			last := mid.next
			mid.next = pre
			if last.next == nil {
				last.next = mid
				break
			}
			pre = mid
			mid = last
		}
	}
	this.head, this.tail = this.tail, this.head
}

func (this *Chain) LeftShift(num int) {
	if num >= this.size {
		this.size = 0
		this.head = nil
		this.tail = nil
	} else {
		temp := this.head
		for i := 0; i < num; i++ {
			temp = temp.next
		}
		this.head = temp
		this.size -= num
	}

}
