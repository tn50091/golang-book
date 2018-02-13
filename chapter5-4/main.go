package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var input int
	var output int = rand.Intn(10)
	fmt.Println("Output: ", output)
	for i := 1; i <= 5; i++ {
		fmt.Print("Enter a number between 1-10: ")
		fmt.Scanf("%d\n", &input)

		//var output int = 5
		if input < output {
			fmt.Println("Input should be greater!!!")
			continue
		} else if input > output {
			fmt.Println("Input should be smaller!!!")
			continue
		} else {
			fmt.Println("Holy ..., U got it!!!!")
			break
		}
	}
}
