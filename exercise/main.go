package main

import (
	"fmt"
)

type VendingMachine struct{
	insertedMoney int
	coins map[string]int
	items map[string]int
}

func (v VendingMachine) InsertedMoney() int {
	return v.insertedMoney
}

func (v *VendingMachine) InsertCoin(coin string) {
	v.insertedMoney += v.coins[coin]
}

func (v *VendingMachine) SelectSD() string {
	price := v.items["SD"]
	change := v.insertedMoney - price
	return "SD"+ v.Change(change)
}

func (v *VendingMachine) SelectCC() string {
	price := v.items["CC"]
	change := v.insertedMoney - price
	return "CC"+ v.Change(change)
}

func (v *VendingMachine) Change(c int) (change string) {
	v.insertedMoney = 0
	var str string
	value := [...]int{10,5,2,1}
	coins := [...]string{"T","F","TW","0"}
	for i:=0;i<len(value);i++ {
		if c >= value[i] { 
			c -= value[i]
			str += ","+coins[i] 
		}
	}
	return str
}

func main() {
	var coins = map[string]int{"T":10,"F":5,"TW":2,"O":1}
	var items = map[string]int{"SD":18,"CC":12,"DW":7}
	vm := VendingMachine{coins: coins, items: items}
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
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())	// Inserted Money: 20
	can = vm.SelectCC()
	fmt.Println("Item/Change:", can)	//CC,F,TW,O
}