package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(weatherCelcius(25, "Bankgok few clouds"))
	fmt.Println(weatherCelcius(34, "Tak sunny"))
	fmt.Println(weatherCelcius(17, "Phuket rainy"))
	fmt.Println(weatherCelcius(9, "Chiang-mai cold"))
	fmt.Println(weatherCelcius(-15, "Very cold"))
}

func weatherCelcius(temp int,desc string) string {
	var out string
	var output [][]string
	str := strconv.Itoa(temp)
	for _,ascii := range str {
		digits:=string(ascii)
		if digits == "1" {
			number1 := []string{"   ", "  |", "  |"}
			output = append(output,number1)
		}else if digits == "2" {
			number2 := []string{" _ ", " _|", "|_ "}
			output = append(output,number2)
		}else if digits == "3" {
			number3 := []string{" _ ", " _|", " _|"}
			output = append(output,number3)
		}else if digits == "4" {
			number4 := []string{"   ", "|_|", "  |"} 
			output = append(output,number4)
		}else if digits == "5" {
			number5 := []string{" _ ", "|_ ", " _|"}
			output = append(output,number5)
		}else if digits == "6" {
			number6 := []string{" _ ", "|_ ", "|_|"}
			output = append(output,number6)
		}else if digits == "7" {
			number7 := []string{" _ ", "  |", "  |"}
			output = append(output,number7)
		}else if digits == "8" {
			number8 := []string{" _ ", "|_|", "|_|"}
			output = append(output,number8)
		}else if digits == "9" {
			number9 := []string{" _ ", "|_|", " _|"}
			output = append(output,number9)
		}else if digits == "-" {
			numberminus := []string{"   ", " _ ", "   "}
			output = append(output,numberminus)
		}else if digits == "." {
			numberdot := []string{"   ", " . ", "   "}
			output = append(output,numberdot)
		}else {
			number0 := []string{" _ ", "| |", "|_|"}
			output = append(output,number0)
		}
	}

	for i:=0;i<3;i++ {
		for j:=0;j<len(str);j++ {
			out+=output[j][i]
		}
		if i == 2 {out+="C"}
		out+="\n"
	}
	out+=desc
	return out

	/*if temp == 25 {
		out1 := " _  _\n"
		out2 := " _||_\n"
		out3 := "|_  _| C\n"
		out4 := desc+"\n"
		output := out1 + out2 + out3 + out4
		return output
	}*/
}