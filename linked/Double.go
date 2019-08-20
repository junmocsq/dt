package linked

import "dt/derror"

type DNode struct {
	Data interface{}
	Next *DNode
	Pre  *DNode
}

type DoubleLinked struct {
	Head   *DNode
	Tail   *DNode
	Length uint32
}

func DCreate() *DoubleLinked {
	return &DoubleLinked{Length: 0, Head: nil, Tail: nil}
}

func DIsEmpty(linked *DoubleLinked) bool {
	if linked.Length == 0 {
		return true
	} else {
		return false
	}
}

func DFindPrevious(data interface{}, linked *DoubleLinked) {

}

func DFind(data interface{}, linked *DoubleLinked) (*DNode, error) {
	if DIsEmpty(linked) {
		return nil, derror.NewErr(4000)
	}
	temp := linked.Head
	for {
		if data == temp.Data {
			return temp, nil
		}
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	return nil, derror.NewErr(4001)
}

func DInsert(data interface{}, linked *DoubleLinked) bool {
	temp := DNode{Data: data}
	if DIsEmpty(linked) {
		linked.Length++
		linked.Head = &temp
		linked.Tail = &temp
		return true
	}
	pre := linked.Tail
	linked.Tail = &temp
	pre.Next = &temp
	linked.Tail.Pre = pre
	linked.Length ++
}

