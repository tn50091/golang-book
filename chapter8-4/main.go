package main

import (
	"fmt"
)

func main() {
	array := [3]int{1,2,3}
	double(&array)
	fmt.Printf("original addr %p\n", &array)
	fmt.Printf("original %v\n", array)
}

func double(number *[3]int) {
	fmt.Printf("double addr %p\n", &number)
	fmt.Printf("double value %v\n", *number)
	for i := range number {
		number[i] *= 2
		//(*number)[i]
	}
	fmt.Println(*number)
}