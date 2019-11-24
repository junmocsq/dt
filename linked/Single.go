package linked

import (
	"dt/derror"
	"fmt"
)

type node struct {
	Data interface{}
	Next *node
}

type singleLinked struct {
	Length uint32
	Head   *node
	Tail   *node
}

func SCreate() *singleLinked {
	head := singleLinked{Length: 0, Head: nil, Tail: nil}
	return &head
}

func SIsEmpty(linked *singleLinked) bool {
	if linked.Length == 0 {
		return true
	} else {
		return false
	}
}

func SIsLast(node *node, linked *singleLinked) bool {
	if linked.Tail == node {
		return true
	} else {
		return false
	}
}

func GetLength(linked *singleLinked) uint32 {
	return linked.Length
}

/**
查找指定数据的上一个节点
*/
func SFindPrevious(data interface{}, linked *singleLinked) (*node, error) {

	if linked.Length == 0 {
		return &node{}, derror.NewErr(4000)
	}
	temp := linked.Head
	if temp.Next == nil {
		return &node{}, derror.NewErr(4001)
	}
	for {
		if temp.Next.Data == data {
			return temp, nil
		}
		if temp.Next.Next == nil {
			return &node{}, derror.NewErr(4001)
		}
		temp = temp.Next
	}
}

func SFind(data interface{}, linked *singleLinked) (*node, error) {
	if linked.Length == 0 {
		return &node{}, derror.NewErr(4000)
	}
	temp := linked.Head
	for {
		if temp.Data == data {
			return temp, nil
		}
		if temp.Next == nil {
			return &node{}, derror.NewErr(4001)
		}
		temp = temp.Next
	}
}

/**
插入队尾 不支持数据重复
*/
func SInsert(data interface{}, linked *singleLinked) bool {
	node := node{Data: data, Next: nil}
	if linked.Length == 0 {
		linked.Head = &node
		linked.Tail = &node
		linked.Length++
		return true
	}
	temp := linked.Head
	for {
		if temp.Data == data {
			return false
		}
		if temp.Next != nil {
			temp = temp.Next
		} else {
			break
		}
	}
	temp.Next = &node
	linked.Tail = &node
	linked.Length++
	return true
}

func SDelete(data interface{}, linked *singleLinked) error {
	if linked.Length == 0 {
		return derror.NewErr(4000)
	}
	// 长度为1 直接判断
	if linked.Length == 1 && linked.Head.Data == data {
		linked.Head = nil
		linked.Tail = nil
		linked.Length--
	}
	// 查找要删除节点的上一个节点
	pre, err := SFindPrevious(data, linked)
	if err != nil {
		return err
	}
	next := pre.Next.Next
	next.Next = nil
	if next == nil {
		linked.Tail = pre
	}
	pre.Next = next
	linked.Length--
	return nil
}

func SDeleteList(linked *singleLinked) {
	p := linked.Head
	linked = nil
	for p.Next != nil {
		temp := p
		p = nil
		p = temp.Next
	}
}

func SPrint(linked *singleLinked) {
	temp := linked.Head
	index := 0
	for temp != nil {
		fmt.Println(index, temp.Data)
		index++
		temp = temp.Next
	}
}
