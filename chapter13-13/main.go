package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	go func() { ch <- 1}()
	go func() { ch <- 2}()
	go func() { ch <- 3}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}