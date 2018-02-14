package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 15; i++ {
		result := FizzBuzz(i)
		fmt.Println(i, result)
	}
}

func FizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return ""
	}
}
