package main

import (
	"fmt"
)

func main() {
	// Long Dec
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// Short Dec
	// Type Interface
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	// Constant
	const w string = "Hello, World"
	//a = "Holy Moly"

	// Multiple Var
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Swap
	v1, v2 := "first", "sec"
	v1, v2 = v2, v1
	fmt.Println(v1)
	fmt.Println(v2)
}

/*
func f() {
	fmt.Println(x)
}*/