package cdt

import (
	"errors"
	"math"
)

// 凸包
// 找到点集中的一个点，做一条垂直于y的正向线(0,1)，求出其他点与x的逆时针夹角【极角】，相同极角的按距离选择点的大小排序，从小到大排序
// 从链表里取出相邻的三个点，计算逆时针夹角，大于180c的就是凸包
//

type hull struct {
	x    float64
	y    float64
	next *hull
	pre  *hull
}

type convexHull struct {
	size int
	head *hull
}

func newhull() *convexHull {
	temp := &hull{}
	temp.next = temp
	temp.pre = temp

	convexhull := &convexHull{size: 0, head: temp}
	return convexhull
}

func (this *convexHull) insert(x, y float64, index int) error {
	if index < 0 && index > this.size {
		return errors.New("链表索引越界")
	}
	pre := this.head
	node := &hull{x: x, y: y, next: nil, pre: nil}

	for i := 0; i < index; i++ {
		pre = pre.next
	}
	prenext := pre.next
	pre.next = node
	node.pre = pre
	node.next = prenext
	prenext.pre = node
	this.size++
	return nil
}

func (this *convexHull) erase(index int) (*hull, error) {
	if index < 0 && index >= this.size {
		return nil, errors.New("链表索引越界")
	}
	pre := this.head
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	delhull := pre.next
	pre.next = pre.next.next
	pre.next.pre = pre
	this.size--
	return delhull, nil
}

func (this *convexHull) get(index int) (*hull, error) {
	if index < 0 && index >= this.size {
		return nil, errors.New("链表索引越界")
	}
	pre := this.head
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	return pre.next, nil
}

type polarNode struct {
	score float64
	x     float64
	y     float64
	next  *polarNode
}

func (this *convexHull) convexhull() {
	// 极角大小排序
	// 相邻三个点计算逆时针夹角，留下来的点即是凸包
}

// 计算极角
func polar(a, b *hull) float64 {
	vecx, vecy := b.x-a.x, b.y-a.y
	if vecx == 0 {
		if vecy > 0 {
			return 90
		} else {
			return 270
		}
	}
	t := vecy / vecx

	ang := math.Atan(t)
	c := ang * 180 / 3.14
	if vecx > 0 {
		if c < 0 {
			c = 360 + c
		}
	} else {
		c = 180 + c
	}

	return c
}

// 计算逆时针夹角是否大于180度 0 =180度 -1 小于180度 1 大于180度
func angle(a, b, c *hull) int {
	x1, y1 := a.x-b.x, a.y-b.y
	x2, y2 := c.x-b.x, c.y-b.y
	t := x1*y2 - x2*y1
	if t == 0 {
		return 0
	} else if t > 0 {
		return -1
	} else {
		return 1
	}
}

// 计算长度平方
func geolengthow(a, b *hull) float64 {
	return math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2)
}
