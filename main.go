package main

import (
	"dt/linked"
)

func main() {
	vec := linked.VCreate(2)

	vec.VAdd(10, 10)
	vec.VAdd(10, 9)
	vec.VAdd(10, 8)
	vec.VAdd(10, 7)
	vec.VAdd(10, 6)
	vec.VAdd(3, 20)
	vec.VAdd(10, 10)
	vec.VAdd(10, 9)
	vec.VAdd(10, 8)
	vec.VAdd(10, 7)
	vec.VAdd(10, 6)
	vec.VAdd(3, 20)
	vec.VAdd(10, 10)
	vec.VAdd(10, 9)
	vec.VAdd(10, 8)
	vec.VAdd(10, 7)
	vec.VAdd(10, 6)
	vec.VAdd(3, 20)

	vec.VPrint()
	vec.VDeduplicate()
	vec.VPrint()
}
