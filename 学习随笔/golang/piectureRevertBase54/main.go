package main

import (
	"os"
	"log"
	"encoding/base64"
	"fmt"
)

func main() {
	file, err := os.Open("picture/KeyStone.png")
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	sourceBuf := make([]byte, 500000)

	n, err := file.Read(sourceBuf)
	if err != nil {
		log.Print(err)
	}
	picStr := base64.StdEncoding.EncodeToString(sourceBuf[:n])
	fmt.Println("base64: " + picStr)

}
