package main

import (
	"fmt"
	"log"
)

type Decorator interface {
	Build(c Car) Car
}

type Car struct {
	bodyType           string
	seats              int
	gpsNavigation      bool
	moonroof           bool
	abs                bool
	touchScreen        bool
	collisionAvoidance bool
	leatherFuniture    bool
}

func (c Car) Build(Car) Car {
	return c
}

type BaseDecorator struct {
	Decorator
}

func (d BaseDecorator) Build(c Car) Car {
	return Car{
		bodyType: "Sedan",
		seats:    5,
	}
}

type SedaDecorator struct {
	Decorator
}

func (d SedaDecorator) Build(c Car) Car {
	return Car{
		bodyType: "Sedan",
		seats:    5,
	}
}

type ModernDecorator struct {
	Decorator
}

func (d ModernDecorator) Build(c Car) Car {
	c = d.Decorator.Build(c)
	c.moonroof = true
	c.abs = true
	c.gpsNavigation = true
	c.moonroof = true
	c.touchScreen = true
	c.collisionAvoidance = true
	return c
}

type LuxuryDecorator struct {
	Decorator
}

func (d LuxuryDecorator) Build(c Car) Car {
	c = d.Decorator.Build(c)
	c.leatherFuniture = true
	return c
}

func Info(c Car) string {
	yesNo := func(v bool) string {
		if v {
			return "yes"
		}
		return "no"
	}
	return fmt.Sprintf("body type = %v, #seats = %v, gps = %v, abs = %v, moonroof = %v, touch screen = %v, collision avoidance = %v, leather funiture = %v",
		c.bodyType, c.seats, yesNo(c.gpsNavigation), yesNo(c.abs), yesNo(c.moonroof), yesNo(c.touchScreen), yesNo(c.collisionAvoidance), yesNo(c.leatherFuniture))
}

func main() {
	simpleCarDeco := SedaDecorator{}
	modernDeco := ModernDecorator{
		Decorator: simpleCarDeco,
	}
	luxuryModernDeco := LuxuryDecorator{
		Decorator: modernDeco,
	}
	luxuryDeco := LuxuryDecorator{
		Decorator: simpleCarDeco,
	}
	log.Printf("build simple sedan car: %v", Info(simpleCarDeco.Build(Car{})))
	log.Printf("build modern sedan car: %v", Info(modernDeco.Build(Car{})))
	log.Printf("build luxury and modern sedan car: %v", Info(luxuryModernDeco.Build(Car{})))
	log.Printf("build luxury sedan car: %v", Info(luxuryDeco.Build(Car{})))
}
