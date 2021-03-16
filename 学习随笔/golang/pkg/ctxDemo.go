package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	context.TODO()
	context.WithCancel(context.TODO())
	context.Background()
	context.WithValue(context.TODO(), "hello", "world")
	context.WithDeadline(context.TODO(), time.Now().Add(1 * time.Second))
	context.WithTimeout(context.Background(), 1 * time.Second)
}

func doCtx1 () {
	ctx, cancel := context.WithCancel(context.TODO())
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("child ctx exit!")
		}
	}()
	//睡眠一秒钟
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(3	 * time.Second)
	fmt.Println("parent ctx exit!")
}