package main

import (
	"strings"
	"fmt"
)

func main() {
	var nodeSelector map[string]string = make(map[string]string)
	hello := " \"hello\""
	fmt.Println(hello)
	nodeSelector["test"] = strings.Trim(strings.Trim(hello, " "), "\"")
	fmt.Println(nodeSelector["test"])
	add := func(a,b int) {
		fmt.Println(a+b)
	}
	add(1,2)
}

