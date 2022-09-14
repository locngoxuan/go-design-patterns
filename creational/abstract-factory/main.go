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

func Info(car Car) string {
	return fmt.Sprintf("brand = %v, body-type = %v, and model = %v", car.Brand(), car.BodyType(), car.Model())
}

type CarFactory interface {
	MakeSUV() Car
	MakeSedan() Car
	MakeElectric() Car
}

type CommonCar struct {
	brand    string
	bodyType string
	model    string
}

func (s CommonCar) Brand() string {
	return s.brand
}

func (s CommonCar) BodyType() string {
	return s.bodyType
}

func (s CommonCar) Model() string {
	return s.model
}

type KiaCar struct {
	CommonCar
}

func (s KiaCar) Brand() string {
	return "Kia"
}

type HyundaiCar struct {
	CommonCar
}

func (s HyundaiCar) Brand() string {
	return "Hyundai"
}

type KiaCarFactory struct{}

func (c KiaCarFactory) MakeSUV() Car {
	return &KiaCar{
		CommonCar: CommonCar{
			bodyType: "SUV",
			model:    "Sorento",
		},
	}
}
func (c KiaCarFactory) MakeSedan() Car {
	return &KiaCar{
		CommonCar: CommonCar{
			bodyType: "Sedan",
			model:    "K5",
		},
	}
}
func (c KiaCarFactory) MakeElectric() Car {
	return &KiaCar{
		CommonCar: CommonCar{
			bodyType: "Electric/Sedan",
			model:    "EV6",
		},
	}
}

type HyundaiCarFactory struct{}

func (c HyundaiCarFactory) MakeSUV() Car {
	return &HyundaiCar{
		CommonCar: CommonCar{
			bodyType: "SUV",
			model:    "Palaside",
		},
	}
}
func (c HyundaiCarFactory) MakeSedan() Car {
	return &HyundaiCar{
		CommonCar: CommonCar{
			bodyType: "Sedan",
			model:    "Sonata",
		},
	}
}
func (c HyundaiCarFactory) MakeElectric() Car {
	return &HyundaiCar{
		CommonCar: CommonCar{
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
func main() {
	hyundai, _ := GetCarFactory("hyundai")
	log.Printf("Make SUV car: %s", Info(hyundai.MakeSUV()))
	log.Printf("Make Sedan car: %s", Info(hyundai.MakeSedan()))
	log.Printf("Make Electric car: %s", Info(hyundai.MakeElectric()))

	_, err := GetCarFactory("ferrari")
	if err != nil {
		log.Printf("faild to get car factory: %v", err)
	}
}
