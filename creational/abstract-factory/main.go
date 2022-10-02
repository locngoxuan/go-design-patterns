package main

import (
	"fmt"
	"log"
	"strings"
)

type Car interface {
	Brand() string
	BodyType() string
	Model() string
}

type CarFactory interface {
	MakeSUV() Car
	MakeSedan() Car
	MakeElectric() Car
}

type BaseCar struct {
	brand    string
	bodyType string
	model    string
}

func (s BaseCar) Brand() string {
	return s.brand
}

func (s BaseCar) BodyType() string {
	return s.bodyType
}

func (s BaseCar) Model() string {
	return s.model
}

type KiaCar struct {
	BaseCar
}

func (s KiaCar) Brand() string {
	return "Kia"
}

type HyundaiCar struct {
	BaseCar
}

func (s HyundaiCar) Brand() string {
	return "Hyundai"
}

type KiaCarFactory struct{}

func (c KiaCarFactory) MakeSUV() Car {
	return &KiaCar{
		BaseCar: BaseCar{
			bodyType: "SUV",
			model:    "Sorento",
		},
	}
}
func (c KiaCarFactory) MakeSedan() Car {
	return &KiaCar{
		BaseCar: BaseCar{
			bodyType: "Sedan",
			model:    "K5",
		},
	}
}
func (c KiaCarFactory) MakeElectric() Car {
	return &KiaCar{
		BaseCar: BaseCar{
			bodyType: "Electric/Sedan",
			model:    "EV6",
		},
	}
}

type HyundaiCarFactory struct{}

func (c HyundaiCarFactory) MakeSUV() Car {
	return &HyundaiCar{
		BaseCar: BaseCar{
			bodyType: "SUV",
			model:    "Palaside",
		},
	}
}
func (c HyundaiCarFactory) MakeSedan() Car {
	return &HyundaiCar{
		BaseCar: BaseCar{
			bodyType: "Sedan",
			model:    "Sonata",
		},
	}
}
func (c HyundaiCarFactory) MakeElectric() Car {
	return &HyundaiCar{
		BaseCar: BaseCar{
			bodyType: "Electric/Sedan",
			model:    "IONIQ",
		},
	}
}

func GetCarFactory(brand string) (CarFactory, error) {
	switch strings.ToLower(brand) {
	case "kia":
		return &KiaCarFactory{}, nil
	case "hyundai":
		return &HyundaiCarFactory{}, nil
	default:
		return nil, fmt.Errorf("not found brand %s", brand)
	}
}

func Info(car Car) string {
	return fmt.Sprintf("brand = %v, body-type = %v, and model = %v", car.Brand(), car.BodyType(), car.Model())
}

func main() {
	kia, _ := GetCarFactory("kia")
	log.Printf("Make SUV car: %s", Info(kia.MakeSUV()))
	log.Printf("Make Sedan car: %s", Info(kia.MakeSedan()))
	log.Printf("Make Electric car: %s", Info(kia.MakeElectric()))

	hyundai, _ := GetCarFactory("hyundai")
	log.Printf("Make SUV car: %s", Info(hyundai.MakeSUV()))
	log.Printf("Make Sedan car: %s", Info(hyundai.MakeSedan()))
	log.Printf("Make Electric car: %s", Info(hyundai.MakeElectric()))
}
