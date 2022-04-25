package main

import (
	"fmt"
	"strings"
	"time"
)

func a2() {
	aa := []string{"csq1", "csq2", "csq3", "csq4", "csq5", "csq6", "csq7", "csq8", "csq9"}

	t := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		strings.Join(aa, "-")
	}
	fmt.Println(time.Now().UnixNano() - t)
	t = time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		s := ""
		for _, k := range aa {
			s += k + ","
		}
	}
	fmt.Println(time.Now().UnixNano() - t)

}

func main() {
	//vec := linked.VCreate(2)
	//
	//vec.VAdd(10, 10)
	//vec.VAdd(10, 9)
	//vec.VAdd(10, 8)
	//vec.VAdd(10, 7)
	//vec.VAdd(10, 6)
	//vec.VAdd(3, 20)
	//vec.VAdd(10, 10)
	//vec.VAdd(10, 9)
	//vec.VAdd(10, 8)
	//vec.VAdd(10, 7)
	//vec.VAdd(10, 6)
	//vec.VAdd(3, 20)
	//vec.VAdd(10, 10)
	//vec.VAdd(10, 9)
	//vec.VAdd(10, 8)
	//vec.VAdd(10, 7)
	//vec.VAdd(10, 6)
	//vec.VAdd(3, 20)

	//vec.VPrint()
	//vec.VDeduplicate()
	//vec.VPrint()
	a2()
}
