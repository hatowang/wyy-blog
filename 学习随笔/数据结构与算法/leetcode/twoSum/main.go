package main

import "fmt"

func main() {
	fmt.Println(10 % 10)
}


func twoSum(nums []int, target int) []int {
	another := make(map[int]int)
	for i := 0; i< len(nums) -1; i++ {
		another[target-nums[i]] = i
		if index, ok := another[nums[i+1]];ok {
			return []int {index, i + 1}
		}
	}
	return nil
}