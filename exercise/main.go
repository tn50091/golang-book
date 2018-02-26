package main

import (
	"fmt"
)

type VendingMachine struct{
	insertedMoney int
}

func (v VendingMachine) InsertedMoney() int {
	return v.insertedMoney
}

func (v *VendingMachine) InsertCoin(coin string) {
	v.insertedMoney = 10
}

func main() {
	vm := VendingMachine{}
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// Inserted Money: 0
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// Inserted Money: 10
}