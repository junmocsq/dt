package cdt

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestChain_Output(t *testing.T) {
	chain := &Chain{size: 0, head: nil}
	var l LinearList
	l = chain
	l.Insert(0, 1)
	l.Insert(0, 2)
	l.Insert(0, 3)
	l.Insert(0, 4)
	l.Insert(0, 5)
	l.Insert(5, 6)
	l.Output()

	l = NewArr(10)
	l.Insert(0, 11)
	l.Insert(0, 12)
	l.Insert(0, 13)
	l.Insert(0, 14)
	l.Insert(0, 15)
	l.Insert(5, 16)
	l.Output()
	chain.Erase(5)
	chain.PushBack(12)
	chain.Output()
	chain.Reverse()
	chain.LeftShift(6)
	chain.Output()
}

func TestDoublyLinkedList_Output(t *testing.T) {
	double := &DoublyLinkedList{size: 0}
	var l LinearList
	l = double
	l.Insert(0, 1)
	l.Insert(0, 2)
	l.Insert(0, 3)
	l.Insert(0, 4)
	l.Insert(0, 5)
	l.Insert(5, 6)
	l.Output()
	l.Insert(0, 11)
	l.Insert(6, 12)
	l.Insert(0, 13)
	l.Insert(0, 14)
	l.Insert(0, 15)
	l.Insert(5, 16)
	l.Output()
	double.Erase(5)
	double.Output()

	testbox := &NewBox{}
	for i := 0; i < 1e3; i++ {
		rand.Seed(int64(i))
		score := rand.Int() % 1000
		testbox.Insert(0, score, "junmo-"+strconv.Itoa(time.Now().Nanosecond()))
	}
	t1 := time.Now().UnixNano()
	testbox.binSort(999, 1)
	t2 := time.Now().UnixNano()
	fmt.Println(t2 - t1)
	testbox2 := &NewBox{}
	for i := 0; i < 1e3; i++ {
		rand.Seed(int64(i))
		score := rand.Int() % 1000
		testbox2.Insert(0, score, "junmo-"+strconv.Itoa(time.Now().Nanosecond()))
	}
	t3 := time.Now().UnixNano()
	testbox2.Box(999)
	t4 := time.Now().UnixNano()
	fmt.Println(t4 - t3)

	a := &hull{}
	b := &hull{x: -1.732, y: -1}
	fmt.Println(polar(a, b))

}
