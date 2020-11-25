package main

import "strings"

func main() {
	s := "dvdf"
	lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	strs := make([]int32, len(s))
	vals := make(map[int32]int)
	count := 0
	for _, val := range s {
		strs = append(strs, val)
	}

	for _, val := range strs {
		if _, ok := vals[val];ok {
			count = len(vals)
			if len(strings.Split(s, string(val))) == 2{
				s = strings.Split(s, string(val))[2]
				if count < lengthOfLongestSubstring(s) {
					count = lengthOfLongestSubstring(s)
				}
			}
		}
		vals[val]=0
		count = len(vals)
	}
	return count
}
