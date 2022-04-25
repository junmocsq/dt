package cdt

// 双链表
import (
	"errors"
	"fmt"
)

type DoubleLinkedNode struct {
	ele  int
	next *DoubleLinkedNode
	pre  *DoubleLinkedNode
}

type DoublyLinkedList struct {
	size int
	head *DoubleLinkedNode
	tail *DoubleLinkedNode
}

/**
// 线性表
type LinearList interface {
	Empty() bool
	Size() int
	Get(int) (int, error)
	IndexOf(int) int
	Erase(index int) (int, error)
	Insert(index int, val int) bool
	Output()
}
*/

func (this *DoublyLinkedList) Empty() bool {
	if this.size == 0 {
		return true
	} else {
		return false
	}
}

func (this *DoublyLinkedList) Size() int {
	return this.size
}
func (this *DoublyLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= this.size {
		return 0, errors.New("双链表越界")
	}
	temp := this.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.ele, nil
}
func (this *DoublyLinkedList) IndexOf(ele int) int {
	res := -1
	temp := this.head
	for i := 0; i < this.size; i++ {
		if temp.ele == ele {
			res = i
			break
		}
	}
	return res
}
func (this *DoublyLinkedList) Erase(index int) (int, error) {
	if index < 0 || index >= this.size {
		return 0, errors.New("双链表越界")
	}
	resEle := 0
	if index == 0 {
		resEle = this.head.ele
		if this.size == 1 {
			this.head = nil
			this.tail = nil
		} else {
			this.head = this.head.next
			this.head.pre = nil
		}

	} else if index == this.size-1 {
		resEle = this.tail.ele
		this.tail = this.tail.pre
		this.tail.next = nil
	} else {
		// 从后往前找
		if index > this.size/2-1 {
			temp := this.tail
			index = this.size - 1 - index
			for i := 0; i < index-1; i++ {
				temp = temp.pre
			}
			resEle = temp.pre.ele
			temp.pre = temp.pre.pre
			temp.pre.next = temp

		} else {
			temp := this.head
			for i := 0; i < index-1; i++ {
				temp = temp.next
			}
			resEle = temp.next.ele
			temp.next = temp.next.next
			temp.next.pre = temp
		}
	}
	this.size--
	return resEle, nil
}

func (this *DoublyLinkedList) Insert(index int, val int) bool {
	if index < 0 || index > this.size {
		return false
	}
	node := &DoubleLinkedNode{ele: val}
	if index == 0 {
		if this.size == 0 {
			this.head = node
			this.tail = node
		} else {
			node.next = this.head
			node.next.pre = node
			this.head = node
		}
	} else {
		if index == this.size {
			this.tail.next = node
			node.pre = this.tail
			this.tail = node
		} else {
			// 从后往前找
			if index > this.size/2 {
				temp := this.tail
				index = this.size - index - 1
				for i := 0; i < index; i++ {
					temp = temp.pre
				}
				node.pre = temp.pre
				node.pre.next = node
				node.next = temp
				node.next.pre = node

			} else {
				temp := this.head
				for i := 0; i < index-1; i++ {
					temp = temp.next
				}
				node.pre = temp
				node.next = temp.next
				node.pre.next = node
				node.next.pre = node
			}
		}
	}
	this.size++
	return true
}

func (this *DoublyLinkedList) Output() {
	fmt.Println("size:", this.size, "\n正排：")
	temp := this.head
	for i := 0; i < this.size; i++ {
		fmt.Printf("%d\t", temp.ele)
		if temp.next != nil {
			temp = temp.next
		}
	}
	fmt.Printf("\n倒排：\n")
	temp = this.tail
	for i := 0; i < this.size; i++ {
		fmt.Printf("%d\t", temp.ele)
		if temp.pre != nil {
			temp = temp.pre
		}
	}
	fmt.Printf("\n")
}
