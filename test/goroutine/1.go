package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("vvvv")
	fmt.Println("vim-go")

	time.Sleep(time.Second)
}
