package cdt

type chapter2 struct{}

// horner 法则
func (c *chapter2)f2_4(coeff []int,x int) int {
	length := len(coeff)
	value := coeff[length-1]
	for i:=length-1;i>=1;i--{
		value = value*x + coeff[i-1]
	}
	return value
}

func (c *chapter2)f2_5(rank []int)([]int,[]int) {
	length := len(rank)
	re := make([]int,length)
	//sortRank := make([]int,length)
	for i:=1;i<length;i++{
		for j:=0;j<i;j++{
			if rank[i]>=rank[j]{
				re[i]++
			}else {
				re[j]++
			}
		}
	}

	// 原地重排
	i := 0
	for i<length {
		if re[i] !=i{
			temp := re[i]
			re[i],re[temp] = re[temp],re[i]
			rank[i],rank[temp] = rank[temp],rank[i]
			continue
		}else {
			i++
		}
	}
	//// 根据排序移动
	//for i:=0;i<length;i++{
	//	sortRank[re[i]] = rank[i]
	//}
	return re,rank
}

/**
选择排序
 */
func (c *chapter2)f2_7selectSort(arr []int) {
	length := len(arr)
	sorted := false
	for i:=length-1;i>1&&!sorted;i--{
		max := 0
		sorted = true
		for j:=1;j<i;j++{
			// 找出最大值
			if arr[j]>arr[max]{
				max = j
			}else {
				sorted=false
			}
		}
		if max!=i{
			arr[max],arr[i] = arr[i],arr[max]
		}
	}
}
