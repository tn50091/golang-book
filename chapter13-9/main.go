package main

import (
	"fmt"
)

func main() {
	array := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go sum(array[:len(array)/2], ch)
	go sum(array[len(array)/2:], ch)
	//x, y, z := <-ch, <-ch, <-ch
	//fmt.Println(x, y, z, x+y) 
	x := <-ch
	fmt.Println(x)
} 

func sum(array []int, ch chan int) {
	sum := 0
	for _, value := range array {
		sum += value
	}
	ch <- sum 
}