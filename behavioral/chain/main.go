package main

import (
	"fmt"
	"log"
)

type Car struct {
	bodyType      string
	seats         int
	abs           bool
	gpsNavigation bool
}

func (c Car) Info() string {
	yesNo := func(b bool) string {
		if b {
			return "yes"
		}
		return "no"
	}
	return fmt.Sprintf("body type = %v, #seats = %v, abs = %v, gps = %v",
		c.bodyType, c.seats, yesNo(c.abs), yesNo(c.gpsNavigation))
}

type Step interface {
	SetNext(s Step)
	Setup(Car) Car
}

func CreateCarKittingProcess(steps ...Step) Step {
	first := steps[0]
	for i := 1; i < len(steps); i++ {
		first.SetNext(steps[i])
	}
	return first
}

type BaseStep struct {
	Next Step
}

func (b *BaseStep) SetNext(s Step) {
	if b.Next == nil {
		b.Next = s
		return
	}
	b.Next.SetNext(s)
}

type BodyStep struct {
	BaseStep
	bodyType string
}

func (b BodyStep) Setup(c Car) Car {
	c.bodyType = b.bodyType
	return b.Next.Setup(c)
}

type SeatStep struct {
	seats int
	BaseStep
}

func (b SeatStep) Setup(c Car) Car {
	c.seats = b.seats
	return b.Next.Setup(c)
}

type ABSStep struct {
	BaseStep
}

func (b ABSStep) Setup(c Car) Car {
	c.abs = true
	if b.Next == nil {
		return c
	}
	return b.Next.Setup(c)
}

type GPSStep struct {
	BaseStep
}

func (b GPSStep) Setup(c Car) Car {
	c.gpsNavigation = true
	if b.Next == nil {
		return c
	}
	return b.Next.Setup(c)
}

func main() {
	process := CreateCarKittingProcess(&BodyStep{
		bodyType: "sedan",
	}, &SeatStep{
		seats: 4,
	}, &ABSStep{}, &GPSStep{})
	log.Printf("build a sedan car: %v", process.Setup(Car{}).Info())

	process = CreateCarKittingProcess(&BodyStep{
		bodyType: "sedan",
	}, &SeatStep{
		seats: 4,
	}, &ABSStep{})
	log.Printf("build a sedan car without gps: %v", process.Setup(Car{}).Info())

	process = CreateCarKittingProcess(&BodyStep{
		bodyType: "SUV",
	}, &SeatStep{
		seats: 7,
	}, &ABSStep{})
	log.Printf("build a SUV car without gps: %v", process.Setup(Car{}).Info())
}
