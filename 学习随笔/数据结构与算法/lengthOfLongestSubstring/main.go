package main

import (
	"fmt"
)

func main() {
	s := " "
	fmt.Println(lengthOfLongestSubstring(s))

}

func lengthOfLongestSubstring(s string) int {
	length := 0
	left :=0
	right :=0
	myMap := make(map[uint8]bool)
	for ;left<=right && right < len(s)-1 ;  {
		myMap[s[right]] = true
		if _, ok := myMap[s[right+1]];ok {
			left++
			right++
		} else {
			fmt.Println("ok", left, right)
			right++
			length++
		}
	}
	fmt.Println(left, right)
	length = len(s[left:right+1])
	return length
}

//func lengthOfLongestSubstring(s string) int {
//	length := 0
//	for  i := 0; i < len(s); i++ {
//		for j := len(s)-1;j>=i;j-- {
//		if !hasRepeat(s[i:j+1]) {
//			if len(s[i:j+1]) > length {
//				length = len(s[i:j+1])
//			}
//		}
//	}
//}
//	return length
//}
//
//func hasRepeat(s string) bool {
//	for  i := 0; i < len(s); i++ {
//		for j := len(s) -1 ; j >= i; j-- {
//			if i != j && s[i] == s[j] {
//				return true
//			}
//		}
//	}
//	return false
//}