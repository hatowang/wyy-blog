package main

import (
	"github.com/streadway/amqp"
	"fmt"
	"context"
	"time"
)

func main() {
	TestContext()
	//conn, err := MqDial(fmt.Sprintf("amqp://%s:%s@%s:%s/", url.QueryEscape("guest"), url.QueryEscape("U6feB$W&MO7YeU#4"), "10.110.64.31", "31359"))
	//fmt.Println(err, conn)
}

func MqDial(url string) (*amqp.Connection, error) {
	fmt.Println("url is ", url)
	return amqp.Dial(url)
}

func TestContext() {
	ctx, cancel := context.WithCancel(context.TODO())
	go func(ctx context.Context) {
		for {
			fmt.Println("hello")
			time.Sleep(3 * time.Second)
		}
	}(ctx)
	time.Sleep(7 * time.Second)
	go cancel()
	fmt.Println("取消喽")
	time.Sleep(7 * time.Second)
}
