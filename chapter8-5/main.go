package main

import (
	"fmt"
)

func main() {
	slice := []int{1,2,3}
	double(slice)
	fmt.Printf("original addr %p\n", slice)
	fmt.Printf("original %v\n", slice)
}

func double(number []int) {
	fmt.Printf("double addr %p\n", number)
	//for i := range number {
	for i:=0; i<len(number); i++ {
		number[i] *= 2
	}
	fmt.Println(number)
}