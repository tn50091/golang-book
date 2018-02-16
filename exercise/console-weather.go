package main

import (
	"fmt"
)

func main() {
	fmt.Println(weatherCelcius(25, "Bankgok few clouds"))
	fmt.Println(weatherCelcius(34, "Tak sunny"))
	fmt.Println(weatherCelcius(17, "Phuket rainy"))
	fmt.Println(weatherCelcius(9, "Chiang-mai cold"))
}


func weatherCelcius(temp int,desc string) string {
	var output string = ""
	if temp == 25 {
		out1 := " _  _\n"
		out2 := " _||_\n"
		out3 := "|_  _| C\n"
		out4 := desc+"\n"
		output := out1 + out2 + out3 + out4
		return output
	}
	return output
}