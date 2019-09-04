package cdt

import "testing"

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
	chain.Output()
}
