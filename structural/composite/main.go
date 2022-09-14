package main

import "log"

type Product interface {
	Price() int
}

type Iphone struct{}

func (i Iphone) Price() int {
	return 999
}

type AirPod struct{}

func (i AirPod) Price() int {
	return 249
}

type Ipad struct{}

func (i Ipad) Price() int {
	return 599
}

type ApplePencil struct{}

func (i ApplePencil) Price() int {
	return 129
}

type RegularBox struct {
	products []Product
}

func (r *RegularBox) Add(p Product) {
	r.products = append(r.products, p)
}

func (r RegularBox) Price() int {
	if r.products == nil || len(r.products) == 0 {
		return 0
	}
	v := int(0)
	for _, p := range r.products {
		v = v + p.Price()
	}
	//adding tax
	return v + int(float32(v)*0.075)
}

type LuxuryBox struct {
	RegularBox
}

func (r LuxuryBox) Price() int {
	v := r.RegularBox.Price()
	//extra cost of luxury
	v += 10
	//adding tax
	return v + int(float32(v)*0.075)
}

func main() {
	regularBox := RegularBox{
		products: make([]Product, 0),
	}
	regularBox.Add(Iphone{})
	regularBox.Add(AirPod{})
	regularBox.Add(Ipad{})
	regularBox.Add(ApplePencil{})
	log.Printf("total price of regular box: %v", regularBox.Price())

	LuxuryBox := LuxuryBox{
		RegularBox: RegularBox{
			products: make([]Product, 0),
		},
	}
	LuxuryBox.Add(Iphone{})
	LuxuryBox.Add(AirPod{})
	LuxuryBox.Add(Ipad{})
	LuxuryBox.Add(ApplePencil{})
	log.Printf("total price of luxury box: %v", LuxuryBox.Price())
}
