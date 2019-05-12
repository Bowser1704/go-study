package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
	ch <- 4
	for  i := range ch{
		fmt.Println(i)
	}
}
