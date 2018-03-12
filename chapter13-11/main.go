package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	go func() { ch <- 1}()
	go func() { ch <- 2}()
	ch <- 3
	fmt.Println("cap:", cap(ch))
	fmt.Println("len:", len(ch))
	fmt.Println(<-ch)
}