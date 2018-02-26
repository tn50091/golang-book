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
	v.insertedMoney += v.coins[coin]
}

func (v *VendingMachine) SelectSD() string {
	return "SD"
}

func (v *VendingMachine) SelectCC() string {
	return "CC"
}

func main() {
	var coins = map[string]int{"T":10,"F":5,"TW":2,"O":1}
	vm := VendingMachine{coins: coins}
	fmt.Println("Initial Money:", vm.InsertedMoney())	// Inserted Money: 0
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// Inserted Money: 18
	can := vm.SelectSD()
	fmt.Println("Item/Change:", can)	//SD
	vm.InsertCoin("T")
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// Inserted Money: 12
	can = vm.SelectCC()
	fmt.Println("Item/Change:", can)	//CC
}