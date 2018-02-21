package main

import (
	"strconv"
	"fmt"
)

func main() {
	for i := 1; i <= 15; i++ {
		result := FizzBuzz(i)
		fmt.Println(i, result)
	}
}

func FizzBuzz(i int) string {
	ln := [3]int{15,3,5}
	ret := [3]string{"FizzBuzz","Fizz","Buzz"}
	for j:=0;j<len(ln);j++ {
		if i%ln[j] == 0 {
			return ret[j]
		}
	}
	return strconv.Itoa(i)
}
