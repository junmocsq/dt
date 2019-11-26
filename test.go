package main

import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
)

type test interface {
	run() string
	say() string
}

type testabc interface {
	test
	abc() string
}

type junmo struct {
}

type lxq struct {
}

func (t junmo) run() string {
	return "junmo run"
}

func (t junmo) say() string {
	return "junmo say"
}

func (t junmo) abc() string {
	return "junmo abc"
}

func (t lxq) run() string {
	return "lxq run"
}

func (t lxq) say() string {
	return "lxq say"
}

func main() {
	var j testabc
	j = junmo{}
	var l test
	l = lxq{}

	fmt.Println(j.run())
	fmt.Println(j.say())
	fmt.Println(j.abc())
	//j.abc()

	l = j
	fmt.Println(l.run())
	fmt.Println(l.say())
	//fmt.Println(l.abc())

	var i interface{}
	i = 10
	var abc string
	if o, ok := i.(string); ok {
		abc = o
	}
	fmt.Println(abc)

	switch i := i.(type) {
	case string:
		fmt.Println(i + " is a string")
	case int:
		fmt.Println(i, "is a int")
	default:
		fmt.Println("未知类型")
	}

	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(2)
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	t := time.Now().Nanosecond()
	ch := make(chan int, 10)
	sum := 0
	ch <- sum
	for m := 0; m < 10000; m++ {
		sum = <-ch
		go func() {
			for k := 0; k < 10; k++ {
				sum += k
			}
			ch <- sum
		}()
	}

	fmt.Println(runtime.NumGoroutine(), <-ch, time.Now().Nanosecond()-t)
	close(ch)

	if comma, ok := <-ch; ok {
		fmt.Println(comma)
	}

	fmt.Println(runtime.NumGoroutine(), <-ch, time.Now().Nanosecond()-t)
	//time.Sleep(2 * time.Second)

	done := make(chan struct{})
	for abc:=0;abc<100;abc++{
		go GenerateIntA(done)
	}
	close(done)


	a := make(chan int, 10)
	for n := 0; n < 10; n++ {
		a <- n
	}

	//for v := range a {
	//	a <- v + 1
	//
	//}
	go func() {
		for v := range a {
			fmt.Println(v)
		}
		//close(a)
	}()
	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)

	c := make(chan int,10)

	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	go func() {
		defer close(c)
		for i:=0;i<10;i++{
			c<-i
		}
	}()

	for v := range c{
		fmt.Println(v," len:",len(c))
	}
	//time.Sleep(3*time.Second)
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
}

func GenerateIntA(done <-chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				fmt.Println("GenerateIntA close")
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateIntB(done <-chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				fmt.Println("GenerateIntB close")
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}
