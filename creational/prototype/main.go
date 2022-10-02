package main

import (
	"fmt"
	"log"
)

type Cloneable interface {
	Info() string
	Clone() Cloneable
}

type Car struct {
	brand     string
	model     string
	seats     int
	discWheel string
}

func (c Car) Info() string {
	return fmt.Sprintf("brand = %v, model = %v, number of seats = %v, disc wheel = %v",
		c.brand, c.model, c.seats, c.discWheel)
}

func (c Car) Clone() Car {
	return Car{
		brand:     c.brand,
		model:     c.model,
		seats:     c.seats,
		discWheel: c.discWheel,
	}
}

func main() {
	kiaK5 := Car{
		brand:     "Kia",
		model:     "K5",
		seats:     5,
		discWheel: "16 inches Rim Dark Grey",
	}
	log.Printf("original car: %v", kiaK5.Info())
	firstCopied := kiaK5.Clone()
	firstCopied.discWheel = "18 inches Rim Machined W/ Black"
	log.Printf("1st car: %v", firstCopied.Info())

	secondCopied := kiaK5.Clone()
	secondCopied.discWheel = "19 inches Alloy W/center"
	log.Printf("2nd car: %v", secondCopied.Info())
}
