package main

import (
	"fmt"
)

func main() {
	fmt.Println("========String===========")
	backticks := `Hello world!,
today's good day.`
	fmt.Println(backticks)

	doubleQuotes := "Hello world!,\ntoday's good day."
	fmt.Println(doubleQuotes)
}