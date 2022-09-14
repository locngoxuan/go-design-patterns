package main

import (
	"fmt"
	"log"
	"strings"
)

type Car interface {
	Brand() string
}

type Kia struct {
}

func (c Kia) Brand() string {
	return "Kia"
}

type Hyundai struct {
}

func (c Hyundai) Brand() string {
	return "Hyundai"
}

func GetACar(brand string) (Car, error) {
	switch strings.ToLower(brand) {
	case "kia":
		return &Kia{}, nil
	case "hyundai":
		return &Hyundai{}, nil
	default:
		return nil, fmt.Errorf("not found brand %v", brand)
	}
}

func main() {
	kia, _ := GetACar("kia")
	log.Printf("car of %s", kia.Brand())

	_, err := GetACar("ferrari")
	if err != nil {
		log.Printf("faild to get a car: %v", err)
	}
}
