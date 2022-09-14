package main

import (
	"fmt"
	"log"
)

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

func (c Car) Info() string {
	yesNo := func(v bool) string {
		if v {
			return "yes"
		}
		return "no"
	}
	return fmt.Sprintf("body type = %v, #seats = %v, gps = %v, abs = %v, moonroof = %v, touch screen = %v, collision avoidance = %v, leather funiture = %v",
		c.bodyType, c.seats, yesNo(c.gpsNavigation), yesNo(c.abs), yesNo(c.moonroof), yesNo(c.touchScreen), yesNo(c.collisionAvoidance), yesNo(c.leatherFuniture))
}

type Decorator interface {
	Build(c Car) Car
}

type SedaDecorator struct {
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
	log.Printf("build simple sedan car: %v", simpleCarDeco.Build(Car{}).Info())
	log.Printf("build modern sedan car: %v", modernDeco.Build(Car{}).Info())
	log.Printf("build luxury and modern sedan car: %v", luxuryModernDeco.Build(Car{}).Info())
	log.Printf("build luxury sedan car: %v", luxuryDeco.Build(Car{}).Info())
}
