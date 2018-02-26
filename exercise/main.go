package main

import (
	"fmt"
)

type VendingMachine struct{
	insertedMoney int
	coins map[string]int
}

func (v VendingMachine) InsertedMoney() int {
	return v.insertedMoney
}

func (v *VendingMachine) InsertCoin(coin string) {
	//v.insertedMoney = 10
	if coin == "T" {v.insertedMoney += 10}
	if coin == "F"  {v.insertedMoney += 5}
	if coin == "TW"  {v.insertedMoney += 2}
	if coin == "O"  {v.insertedMoney += 1}
}

func main() {
	vm := VendingMachine{}
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// Inserted Money: 0
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// Inserted Money: 10
}