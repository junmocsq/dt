package main

import (
	 "dt/linked"
)

func main()  {
	link := linked.SCreate()
	linked.SInsert(5,link)
	linked.SInsert(1,link)
	linked.SInsert(2,link)
	linked.SInsert(3,link)
	linked.SInsert(4,link)
	linked.SInsert(4,link)
	linked.SInsert(4,link)
	linked.SInsert(4,link)
	linked.SPrint(link)

}