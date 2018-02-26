package main

import (
	"fmt"
)

type NewVendingMachine struct{
	coins int
	items int
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

func (v *NewVendingMachine) SelectCC() string {
	v.items = 12
	cash := v.coins-v.items
	output := "CC"
	return GenOutput(cash, output)
}

func (v *NewVendingMachine) SelectSD() string {
	v.items = 18
	cash := v.coins-v.items
	output := "SD"
	return GenOutput(cash, output)
}

func (v *NewVendingMachine) CoinReturn() string {
	v.items = 0
	cash := v.coins-v.items
	output := ""
	return GenOutput(cash, output)
}

func GenOutput (cash int, output string) string {
	for cash>0 {
		if cash >= 10 {
			cash -= 10
			output += ",T"
			continue
		}
		if cash >= 5 {
			cash -= 5
			output += ",F"
			continue
		}
		if cash >= 2 {
			cash -= 2
			output += ",TW"
			continue
		}
		if cash >= 1 {
			cash -= 1
			output += ",O"
			continue
		}
	}
	return output
}

func Imain() {
	vm := NewVendingMachine{0,0}
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney())
	//can := vm.SelectCC()
	//can := vm.SelectSD()
	can := vm.CoinReturn()
	fmt.Println("Item,Change:", can)
}
