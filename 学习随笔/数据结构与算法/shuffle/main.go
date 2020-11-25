package main

import (
	"math/rand"
	"fmt"
	"errors"
)
var endpoints = []string {
	"10.110.54.31:3232",
	"10.110.54:3232",
	"10.110.54.42:3232",
	"10.110.54.81:3232",
	"10.110.54.11:3232",
	"10.110.54.113:3232",
	"10.110.54.101:3232",
}

//洗牌算法，用于负载均衡
func shuffle(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

func main()  {
	slice := []int{0,1,2,3,4,5,6}
	shuffle1(slice)
	fmt.Println(slice)
	//若请求api失败则重试
	for i := 0; i < 3; i++ {
		sed := rand.Intn(len(slice))
		fmt.Println(sed)
		err := apiRequest(sed)
		fmt.Println(err)
		if err == nil {
			break
		}
	}
}

func apiRequest(num int) error {
	if num %2 == 0 {
		fmt.Println("请求一次api")
		return nil
	}
	return errors.New("num is not invalid")
}

func shuffle1(indexes  []int) {
	for i := len(indexes); i > 0; i-- {
		lastIdx := i -1
		sed := rand.Intn(i)
		indexes[lastIdx], indexes[sed] = indexes[sed], indexes[lastIdx]
	}
}