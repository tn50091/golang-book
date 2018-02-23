package main

import (
	"fmt"
)

type NewVendingMachine struct{
	coins int
	items string
}

func (v *NewVendingMachine) InsertCoin(text string) {
	if text == "T" {v.coins += 10}
	if text == "F"  {v.coins += 5}
	if text == "TW"  {v.coins += 2}
	if text == "O"  {v.coins += 1}
}

func (v NewVendingMachine) GetInsertedMoney() int{
	return v.coins
}

func (v NewVendingMachine) SelectCC() int {
	//v.items = "CC"
	//if v.coins > 12 {return v.items,v.coins-12}
	//return v.items,v.coins
	if v.coins > 12 {return v.coins-12}
	return v.coins
}

func main() {
	vm := NewVendingMachine{0,""}
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 20
	can := vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O
}
