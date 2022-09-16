package main

import (
	"log"
)

type Airplane interface {
	PrepareForLanding()
	Landing()
	PermitLanding()
}

type Mediator interface {
	CanLanding(Airplane) bool
	NotifyAboutLanding()
}

type AmericanAirplane struct {
	mediator Mediator
}

func (a *AmericanAirplane) Landing() {
	log.Println("AmericanAirplane: Landed")
	a.mediator.NotifyAboutLanding()
}
func (a *AmericanAirplane) PrepareForLanding() {
	if !a.mediator.CanLanding(a) {
		log.Println("AmericanAirplane ask for landing: Runway is blocked, waiting")
		return
	}
	log.Println("AmericanAirplane ask for landing: accepted")
}
func (a *AmericanAirplane) PermitLanding() {
	log.Println("AmericanAirplane receive notification: Runaway is clear")
}

type UnitedAirplane struct {
	mediator Mediator
}

func (a *UnitedAirplane) Landing() {
	log.Println("UnitedAirplane: Landed")
	a.mediator.NotifyAboutLanding()
}
func (a *UnitedAirplane) PrepareForLanding() {
	if !a.mediator.CanLanding(a) {
		log.Println("UnitedAirplane ask for landing: Runway is blocked, waiting")
		return
	}
	log.Println("UnitedAirplane ask for landing: accepted")
}
func (a *UnitedAirplane) PermitLanding() {
	log.Println("UnitedAirplane receive notification: Runaway is clear")
}

type DeltaAirplane struct {
	mediator Mediator
}

func (a *DeltaAirplane) Landing() {
	log.Println("DeltaAirplane: Landed")
	a.mediator.NotifyAboutLanding()
}
func (a *DeltaAirplane) PrepareForLanding() {
	if !a.mediator.CanLanding(a) {
		log.Println("DeltaAirplane ask for landing: Runway is blocked, waiting")
		return
	}
	log.Println("DeltaAirplane ask for landing: accepted")
}
func (a *DeltaAirplane) PermitLanding() {
	log.Println("DeltaAirplane receive notification: Runaway is clear")
}

type Airport struct {
	isRunwayFree bool
	landingQueue []Airplane
}

func (a *Airport) CanLanding(airplane Airplane) bool {
	if a.isRunwayFree {
		a.isRunwayFree = false
		return true
	}
	a.landingQueue = append(a.landingQueue, airplane)
	return false
}

func (a *Airport) NotifyAboutLanding() {
	if !a.isRunwayFree {
		a.isRunwayFree = true
	}
	if len(a.landingQueue) > 0 {
		for _, airplane := range a.landingQueue {
			airplane.PermitLanding()
		}
		a.landingQueue = make([]Airplane, 0)
	}
}

func main() {
	newark := Airport{
		landingQueue: make([]Airplane, 0),
		isRunwayFree: true,
	}
	americanAirplane := &AmericanAirplane{
		mediator: &newark,
	}
	deltaAirplane := &DeltaAirplane{
		mediator: &newark,
	}
	unitedAirplane := &UnitedAirplane{
		mediator: &newark,
	}
	log.Println("> 10:05 AM")
	americanAirplane.PrepareForLanding()
	deltaAirplane.PrepareForLanding()
	unitedAirplane.PrepareForLanding()
	log.Println("> 10:15 AM")
	americanAirplane.Landing()
	deltaAirplane.PrepareForLanding()
	unitedAirplane.PrepareForLanding()
	log.Println("> 10:25 AM")
	deltaAirplane.Landing()
	unitedAirplane.PrepareForLanding()
	log.Println("> 10:35 AM")
	unitedAirplane.Landing()
}
