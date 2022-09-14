package main

import "log"

type USLaptop struct {
}

func (l USLaptop) Charge(s Socket) {
	log.Printf("Step to charge laptop with %v socket", s.Name())
	s.PlugThePlugIntoSocket()
}

type Socket interface {
	Name() string
	PlugThePlugIntoSocket()
}

type USSocket struct{}

func (u USSocket) Name() string {
	return "US"
}
func (u USSocket) PlugThePlugIntoSocket() {
	log.Println("- laptop has been plugged into socket")
}

type EuroSocket struct {
	Source string
}

func (u EuroSocket) Name() string {
	return "Euro"
}
func (u EuroSocket) PlugThePlugIntoSocket() {
	log.Printf(" - %s has been plugged into socket", u.Source)
}

type EuroSocketAdapter struct {
	EuroSocket
}

func (u EuroSocketAdapter) PlugThePlugIntoSocket() {
	log.Println(" - laptop has been plugged into adapter")
	u.EuroSocket.Source = "adapter"
	u.EuroSocket.PlugThePlugIntoSocket()
}

func main() {
	laptop := USLaptop{}
	log.Println(">>>>>> staying in US >>>>>>")
	laptop.Charge(USSocket{})
	log.Println(">>>>>> flighting to Euro >>>>>>")
	laptop.Charge(EuroSocketAdapter{
		EuroSocket: EuroSocket{},
	})
}
