package main

import (
	"log"
	"time"
)

type MachineState int

const (
	WaitForSelect = iota
	WaitForPayment
	WaitForPickUp
)

type VendingMachine struct {
	state MachineState
}

func (v *VendingMachine) changeState(s MachineState) {
	v.state = s
}

func (v *VendingMachine) SelectItem() {
	if v.state == WaitForPickUp {
		log.Printf("Machine: invalid action")
		return
	}
	v.changeState(WaitForPayment)
}

func (v *VendingMachine) ProcessPayment() {
	if v.state != WaitForPayment {
		log.Printf("Machine: invalid action")
		return
	}
	v.changeState(WaitForPickUp)
	v.DropSelectedItem()
}

func (v *VendingMachine) DropSelectedItem() {
	log.Printf("Machine: drop selected items")
	<-time.After(1 * time.Second)
	v.changeState(WaitForSelect)
}

type Customer struct {
	*VendingMachine
}

func (c Customer) PressButtonForSelectingItem() {
	log.Printf("Customer press button for selecting item")
	c.VendingMachine.SelectItem()
}

func (c Customer) InsertMoneyForPayment() {
	log.Printf("Customer insert amount of money for payment")
	c.VendingMachine.ProcessPayment()
}

func (c Customer) Pickup() {
	log.Printf("Customer pickup items")
}

func main() {
	vendingMachine := &VendingMachine{
		state: WaitForSelect,
	}
	log.Printf(">>> Customer 01")
	c1 := Customer{
		VendingMachine: vendingMachine,
	}
	c1.PressButtonForSelectingItem()
	c1.PressButtonForSelectingItem()
	c1.InsertMoneyForPayment()
	c1.Pickup()

	log.Printf(">>> Customer 02")
	c2 := Customer{
		VendingMachine: vendingMachine,
	}
	c2.InsertMoneyForPayment()
	c2.PressButtonForSelectingItem()
	c2.InsertMoneyForPayment()
	c2.Pickup()

}
