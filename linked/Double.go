package linked

import "dt/derror"

type dNode struct {
	Data interface{}
	Next *dNode
	Pre  *dNode
}

type doubleLinked struct {
	Head   *dNode
	Tail   *dNode
	Length uint32
}

func DCreate() *doubleLinked {
	return &doubleLinked{Length: 0, Head: nil, Tail: nil}
}

func DIsEmpty(linked *doubleLinked) bool {
	if linked.Length == 0 {
		return true
	} else {
		return false
	}
}

func DFindPrevious(data interface{}, linked *doubleLinked) {

}

func DFind(data interface{}, linked *doubleLinked) (*dNode, error) {
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

func DInsert(data interface{}, linked *doubleLinked) bool {
	temp := dNode{Data: data}
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

