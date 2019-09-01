package cdt

import (
	"fmt"
	"testing"
)
var c  = &chapter2{}
func TestF2_4(t *testing.T)  {
	coeff := []int{1,0,3,4,5}
	x := 10

	fmt.Println(c.f2_4(coeff,x)) //54301
}

func TestF2_5(t *testing.T)  {
	a := []int{10,10,4,1,3}
	fmt.Println(c.f2_5(a)) //[3 4 2 0 1]
}

func TestF2_7(t *testing.T)  {
	a := []int{2,171,4,5,132,34,12,55,98}
	c.f2_7selectSort(a)
	fmt.Println(a)
}