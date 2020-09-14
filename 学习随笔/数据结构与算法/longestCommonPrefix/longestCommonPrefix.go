package main

import (
	"strings"
	"fmt"
)

func main() {
	strs := []string {
		"aaa",
		"aa",
		"aaa",
	}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return  ""
	}
	var prefix  = strs[0]
	for _, k := range strs {
		if len(k) == 0 {
			return  ""
		}
		for !strings.HasPrefix(k, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}