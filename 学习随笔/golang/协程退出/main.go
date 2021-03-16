package main

import (
	"time"
	"fmt"
)

func main() {
	//stop := make(chan struct{})
	//
	//go func() {
	//		time.Sleep(5 * time.Second)
	//		stop
	//}()
	//
	//defer close(stop)
	//
	//select {
	//
	//}
	go DoTest()
	select {

	}
}



func DoTest() {
	go func() {
		fmt.Println("sleep 3s")
		time.Sleep(3*time.Second)
		panic("")
	}()
	fmt.Println("主函数退出了")

}