package main

import (
	"log"
	"sync"
)

type Car struct{}

type CarFactory struct {
	totalCar int
}

func (c *CarFactory) MakeCar() Car {
	c.totalCar++
	return Car{}
}

func (c CarFactory) GetTotalCar() int {
	return c.totalCar
}

var factory *CarFactory

func init() {
	factory = &CarFactory{
		totalCar: 0,
	}
}

func getCarFactory() *CarFactory {
	return factory
}

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			f := getCarFactory()
			f.MakeCar()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("total car has been made: %v", getCarFactory().GetTotalCar())
}
