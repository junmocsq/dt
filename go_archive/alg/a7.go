package alg

import "fmt"

type Alg7 struct {
}

func (this *Alg7) QuickSort(arr []int, start, end int) {
	if end-start < 2 {
		return
	}
	//mid := (start + end - 1) / 2
	//fmt.Println(arr, start, mid, end)
	//for i := start; i < end; {
	//	// 发生替换后 不能向前走，需要再比对一次
	//	if arr[i] > arr[mid] && i < mid {
	//		temp := arr[i]
	//		for j := i; j < mid; j++ {
	//			arr[j] = arr[j+1]
	//		}
	//		arr[mid] = temp
	//		mid--
	//		// = 号很关键 可以排序数字相同的序列
	//	} else if arr[i] <= arr[mid] && i > mid {
	//		temp := arr[i]
	//		for j := i; j > mid; j-- {
	//			arr[j] = arr[j-1]
	//		}
	//		arr[mid] = temp
	//		mid++
	//	} else {
	//		i++
	//	}
	//}

	ntemp := arr[end-1]
	mid := start - 1
	//fmt.Println(arr, start, mid, end)
	for i := start; i < end-1; i++ {
		if arr[i] >= ntemp {
			mid++
			if mid != i {
				arr[i], arr[mid] = arr[mid], arr[i]
			}
		}
	}
	mid++
	arr[end-1], arr[mid] = arr[mid], arr[end-1]

	fmt.Println(arr, start, mid, end)
	this.QuickSort(arr, start, mid)
	this.QuickSort(arr, mid+1, end)
}
