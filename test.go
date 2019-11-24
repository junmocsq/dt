package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {

	//a := make(chan int, 1024)
	//b := make(chan int, 1024)
	//for i := 0; i < 100; i++ {
	//	fmt.Printf("第%d次，", i)
	//	a <- 1
	//	b <- 1
	//	select {
	//	case <-a:
	//		fmt.Println("from a")
	//	case <-b:
	//		fmt.Println("from b")
	//	}
	//}
	ch := make(chan int, 1)
	defer close(ch)
	for i := 0; i < 10; i++ {
		go safeJunmo(ch)
	}
	go func() {
		for i := 0; i < 100000; i++ {
			ch <- i
		}
	}()
	fmt.Println([]byte("junmo"))
	time.Sleep(12 * time.Second)

}

func safeJunmo(ch chan int) {
	fmt.Println("test junmo start")
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error %s\n", err)
		}
	}()
	junmo(ch)
	fmt.Println("test junmo end")
}
func junmo(ch chan int) {
	timeout := make(chan int, 1)
	defer close(timeout)
	go func() {
		time.Sleep(3 * time.Second)
		timeout <- 1
	}()
LABLE:
	for {
		select {
		case _ = <-ch:
			//fmt.Println("ch:",i)
		case <-timeout:
			panic("test error")
			fmt.Println("超时啦")
			break LABLE
		}
	}

}

func (t *BinTree) delete(ele interface{}, node, parent *BinTreeNode) error {
	res, err := t.Compare(ele, node.Elem)
	if err != nil {
		return err
	}
	if node == parent {
		if res == 0 {
			// 相等
			if node.Left != nil && node.Right != nil {
				temp, err := t.FindMin(node.Right)
				if err != nil {
					return err
				}
				node.Elem = temp.Elem
				return t.delete(temp.Elem, node.Right, node)
			} else {
				if node.Right == nil {
					t.Root = node.Left
				} else {
					t.Root = node.Right
				}
			}
		} else if res == 1 {
			return t.delete(ele, node.Right, node)
		} else {
			return t.delete(ele, node.Left, node)
		}
	}

	if res == 1 {
		if node.Right == nil {
			return errors.New("不存在要删除的元素")
		}
		return t.delete(ele, node.Right, node)
	} else if res == -1 {
		if node.Left == nil {
			return errors.New("不存在要删除的元素")
		}
		return t.delete(ele, node.Left, node)
	} else {
		// 相等
		if node.Left != nil && node.Right != nil {
			temp, err := t.FindMin(node.Right)
			if err != nil {
				return err
			}
			node.Elem = temp.Elem
			return t.delete(temp.Elem, node.Right, node)
		} else {
			if parent.Left == node {
				if node.Right == nil {
					parent.Left = node.Left
				} else {
					parent.Left = node.Right
				}
			} else {
				if node.Right == nil {
					parent.Right = node.Left
				} else {
					parent.Right = node.Right
				}
			}
		}
	}
	return nil
}
