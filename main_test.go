package main

import (
	"testing"
	"net"
	"time"
	"fmt"
)

func Array() (x [1024]int) {
	for i := 0; i < 1024; i++ {
		x[i] = i
	}
	return x
}

func Slice() (y []int) {
	y = make([]int, 1024)
	for i :=0; i < 1024; i++ {
		y[i] = i
	}
	return y
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for y := 0; y < b.N; y++ {
		Slice()
	}
}

func TestTime_RawValue(t *testing.T) {
	conn, _ := net.Dial("tcp", ":8080")
	conn.Write([]byte("123"))
	time.Sleep(time.Second * 4)
	rBytes := make([]byte, 0)
	conn.Read(rBytes)
	fmt.Println(string(rBytes))
}