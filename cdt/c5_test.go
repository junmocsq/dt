package cdt

import (
	"fmt"
	"testing"
)

func TestF5_ArrLinearList(t *testing.T) {
	newArr := NewArr(4)
	newArr.Insert(0,1)
	newArr.Insert(0,2)
	newArr.Insert(0,3)
	newArr.Insert(0,4)
	newArr.Output()
	newArr.Insert(0,5)
	newArr.Insert(0,6)
	newArr.Insert(0,7)
	newArr.Insert(0,8)
	newArr.Insert(0,9)
	newArr.Output()

	v,err := newArr.Get(5)
	if err==nil{
		fmt.Printf("index:5 val:%d \n",v)
	}

	fmt.Println("is empty:",newArr.Empty()," size:",newArr.Size()," array[index] = 9 index:",newArr.IndexOf(9))
	newArr.Erase(0)
	newArr.Erase(0)
	newArr.Erase(0)
	newArr.Erase(0)
	newArr.Erase(0)
	newArr.Output()

}
