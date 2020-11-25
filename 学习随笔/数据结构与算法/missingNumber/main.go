package  main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0,1,3}))
}

func missingNumber(nums []int) int {
	n:= len(nums)
	for i:=0; i < n; i++ {
		n ^= nums[i] ^ i
	}

	return n
}