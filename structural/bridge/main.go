package main

import (
	"log"
)

type Color interface {
	GetColor() string
}

type Red struct {
}

func (c Red) GetColor() string {
	return "red"
}

type Blue struct {
}

func (c Blue) GetColor() string {
	return "blue"
}

type Shape interface {
	SetColor(c Color)
	Draw()
}

type CommonShape struct {
	c Color
}

func (s *CommonShape) SetColor(c Color) {
	s.c = c
}

func (s *CommonShape) Draw() {

}

type Rectangle struct {
	CommonShape
}

func (s *Rectangle) Draw() {
	log.Printf("draw rectange with color = %v", s.c.GetColor())
}

type Square struct {
	CommonShape
}

func (s *Square) Draw() {
	log.Printf("draw square with color = %v", s.c.GetColor())
}

type Circle struct {
	CommonShape
}

func (s *Circle) Draw() {
	log.Printf("draw circle with color = %v", s.c.GetColor())
}

func main() {
	rec := Rectangle{}
	rec.SetColor(Red{})
	rec.Draw()
	rec.SetColor(Blue{})
	rec.Draw()
	log.Println("")
	circle := Circle{}
	circle.SetColor(Red{})
	circle.Draw()
	circle.SetColor(Blue{})
	circle.Draw()
	log.Println("")
	square := Square{}
	square.SetColor(Red{})
	square.Draw()
	square.SetColor(Blue{})
	square.Draw()
}
