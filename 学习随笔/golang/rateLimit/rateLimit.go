package main

import (
	"golang.org/x/time/rate"
	"context"
	"fmt"
	"time"
)

func main() {
	for {
		var count = 0
		count = count + 1
		ctx, _ := context.WithCancel(context.TODO())
		go TestLimit(ctx, count)

	}

}


func TestLimit(ctx context.Context, id int) {
	var limit = rate.NewLimiter(1, 1)
	if  limit.AllowN(time.Now(), 1) {
		fmt.Println("token afford！", id)
	}

	fmt.Println("false!", id, limit.Burst())
}