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

func (v *NewVendingMachine) SelectCC() string {
	v.items = "CC"
	cash := v.coins-12
	output := v.items
	return GenOutput(cash, output)
}

func (v *NewVendingMachine) SelectSD() string {
	v.items = "SD"
	cash := v.coins-18
	output := v.items
	return GenOutput(cash, output)
}

func (v *NewVendingMachine) CoinReturn() string {
	v.items = ""
	cash := v.coins-0
	output := v.items
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

func main() {
	vm := NewVendingMachine{0,""}
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
