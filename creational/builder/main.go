package main

import (
	"fmt"
	"log"
)

type Transmission int

const (
	AUTO Transmission = 0
	MANUAL
)

type Car struct {
	bodyType           string
	seats              int
	transmission       Transmission
	gps                bool
	moonroof           bool
	abs                bool
	touchScreen        bool
	collisionAvoidance bool
}

type Builder interface {
	SetBodyType()
	SetSeats()
	SetTransmission()
	SetGPS()
	SetMoonroof()
	SetABS()
	SetTouchSCreen()
	SetCollisionAvoidance()
	GetCar() Car
}
type Director struct {
}

func (d Director) MakeACar(builder Builder) Car {
	builder.SetBodyType()
	builder.SetSeats()
	builder.SetMoonroof()
	builder.SetABS()
	builder.SetGPS()
	builder.SetCollisionAvoidance()
	builder.SetTouchSCreen()
	builder.SetTransmission()
	return builder.GetCar()
}

type CarBuilder struct {
	bodyType           string
	seats              int
	transmission       Transmission
	gps                bool
	moonroof           bool
	abs                bool
	touchScreen        bool
	collisionAvoidance bool
}

func (s *CarBuilder) SetSeats()              {}
func (s *CarBuilder) SetTransmission()       {}
func (s *CarBuilder) SetBodyType()           {}
func (s *CarBuilder) SetGPS()                {}
func (s *CarBuilder) SetMoonroof()           {}
func (s *CarBuilder) SetABS()                {}
func (s *CarBuilder) SetTouchSCreen()        {}
func (s *CarBuilder) SetCollisionAvoidance() {}

func (c CarBuilder) GetCar() Car {
	return Car{
		bodyType:           c.bodyType,
		seats:              c.seats,
		transmission:       c.transmission,
		gps:                c.gps,
		moonroof:           c.moonroof,
		abs:                c.abs,
		touchScreen:        c.touchScreen,
		collisionAvoidance: c.collisionAvoidance,
	}
}

type SimpleSedanCarBuilder struct {
	CarBuilder
}

func (s *SimpleSedanCarBuilder) SetSeats() {
	s.seats = 4
}

func (s *SimpleSedanCarBuilder) SetTransmission() {
	s.transmission = MANUAL
}

func (s *SimpleSedanCarBuilder) SetBodyType() {
	s.bodyType = "Sedan"
}

type SimpleSUVCar struct {
	CarBuilder
}

func (s *SimpleSUVCar) SetSeats() {
	s.seats = 7
}

func (s *SimpleSUVCar) SetTransmission() {
	s.transmission = MANUAL
}

func (s *SimpleSUVCar) SetBodyType() {
	s.bodyType = "SUV"
}

type MidLevelSUVBuilder struct {
	SimpleSUVCar
}

func (s *MidLevelSUVBuilder) SetTransmission() {
	s.transmission = AUTO
}
func (s *MidLevelSUVBuilder) SetGPS() {
	s.gps = true
}
func (s *MidLevelSUVBuilder) SetABS() {
	s.abs = true
}
func (s *MidLevelSUVBuilder) SetMoonroof() {
	s.moonroof = true
}

func (c Car) Info() string {
	getTransimssion := func(t Transmission) string {
		if t == MANUAL {
			return "manual"
		}
		return "auto"
	}
	trueFalse := func(b bool) string {
		if b {
			return "yes"
		}
		return "no"
	}
	return fmt.Sprintf("body type = %v, number of seats = %v, transimission = %v, gps = %v, abs = %v, moonroof = %v",
		c.bodyType, c.seats, getTransimssion(c.transmission), trueFalse(c.gps), trueFalse(c.abs), trueFalse(c.moonroof))
}

func main() {
	d := Director{}
	log.Printf("simple Sedan car: %v", d.MakeACar(&SimpleSedanCarBuilder{}).Info())
	log.Printf("simple SUV car: %v", d.MakeACar(&SimpleSUVCar{}).Info())
	log.Printf("mid SUV car: %v", d.MakeACar(&MidLevelSUVBuilder{}).Info())
}
