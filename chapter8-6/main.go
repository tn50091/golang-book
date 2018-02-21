package main

import (
	"fmt"
)

func main() {
	m := map[int]int{1:1,2:2,3:3}
	double(m)
	fmt.Printf("original addr %p\n", m)
	fmt.Printf("original %v\n", m)
}

func double(number map[int]int) {
	fmt.Printf("double addr %p\n", number)
	for i := range number {
		number[i] *= 2
		//(*number)[i]
	}
	fmt.Println(number)
}