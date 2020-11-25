package main

import (
	"fmt"
)

func main() {
	arr := []int{4,9,2,5,7,1,8}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("the arr is ", arr)
}

func quickSort(arr []int, left, right int) {
	fmt.Println(arr)
	if left >= right {
		return
	}
	temp := arr[left] //定义标志位
	//设置哨兵
	i, j := left, right

	for {
		//找到右边比temp小的数据
		for arr[j] >= temp && j > i  {
			j--
		}
		//找到左边比temp大的数据
		for arr[i] < temp && i < j {
			i++
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
	quickSort(arr, left, i)
	quickSort(arr, i+1 , right)
}