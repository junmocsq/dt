package cdt

import (
	"errors"
	"fmt"
)

// 箱子排序
type studentRecord struct {
	score int
	name  string
	next  *studentRecord
}

type box struct {
	studentNext *studentRecord
}

type NewBox struct {
	size int
	head *studentRecord
}

func (this *NewBox) Empty() bool {
	if this.size == 0 {
		return true
	} else {
		return false
	}
}

func (this *NewBox) Size() int {
	return this.size
}

func (this *NewBox) Get(index int) (*studentRecord, error) {
	if this.size <= index || index < 0 {
		return nil, errors.New("访问的元素越界")
	}
	temp := this.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp, nil
}

func (this *NewBox) IndexOf(val string) int {
	temp := this.head
	for i := 0; i < this.size; i++ {
		if temp.name == val {
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

func (this *NewBox) Erase(index int) (*studentRecord, error) {
	if this.size <= index || index < 0 {
		return nil, errors.New("访问的元素越界")
	}
	var val *studentRecord
	if index == 0 {
		val = this.head
		this.head = this.head.next
	} else {
		temp := this.head
		for i := 0; i < index-1; i++ {
			temp = temp.next
		}
		val = temp.next
		temp.next = temp.next.next

	}
	this.size--
	return val, nil

}

func (this *NewBox) Insert(index int, score int, name string) bool {
	if index > this.size || index < 0 {
		return false
	}
	node := studentRecord{score: score, name: name, next: nil}
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
	this.size++
	return true
}

func (this *NewBox) Output() {
	fmt.Println("box size:", this.size)
	temp := this.head
	for i := 0; i < this.size; i++ {
		fmt.Printf("name:%s score:%d \t", temp.name, temp.score)
		if temp.next != nil {
			temp = temp.next
		}
	}
	fmt.Printf("\n")
}

func newBox(length int) []*box {
	temp := make([]*box, length)
	return temp
}

func boxinsert(b []*box, i int, record *studentRecord) {
	record.next = nil

	temp := b[i]
	if temp == nil {
		b[i] = &box{studentNext: record}
	} else {
		first := temp.studentNext
		for {
			if first.next == nil {
				break
			} else {
				first = first.next
			}
		}
		first.next = record
	}
}

// 箱子排序
func (this *NewBox) Box(totalScore int) {
	box := newBox(totalScore + 1)
	size := this.size
	for i := 0; i < size; i++ {
		record, err := this.Erase(0)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(record,i,this.size)
		boxinsert(box, record.score, record)
	}

	this.head = nil
	this.size = 0

	for i := 0; i < totalScore+1; i++ {
		if box[i] != nil {
			trecord := box[i].studentNext
			for {
				this.Insert(0, trecord.score, trecord.name)
				if trecord.next == nil {
					break
				}
				trecord = trecord.next
			}
		}
	}

}

func (this *NewBox) binSort(maxNum int) {
	//tempNum := 10
	//maxNum*=10
	//for ;maxNum>0;maxNum = maxNum/10 {
	//
	//}
}
