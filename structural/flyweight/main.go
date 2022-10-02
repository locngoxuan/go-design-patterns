package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Point struct {
	PointType
	x int
	y int
}

type PointType struct {
	Color string
	Char  string
}

func (p PointType) GetKey() string {
	return fmt.Sprintf("%s%s", p.Char, p.Color)
}

type PointTypeFactory struct {
	pointTypes map[string]PointType
}

func (p *PointTypeFactory) getPointType(typ string) (PointType, error) {
	v, ok := p.pointTypes[typ]
	if ok {
		return v, nil
	}
	switch typ {
	case "xred":
		v = PointType{
			Color: "red",
			Char:  "x",
		}
	case "ored":
		v = PointType{
			Color: "red",
			Char:  "o",
		}
	case "xblue":
		v = PointType{
			Color: "blue",
			Char:  "x",
		}
	case "oblue":
		v = PointType{
			Color: "blue",
			Char:  "o",
		}
	default:
		return v, fmt.Errorf("wrong point type passed")
	}
	p.pointTypes[typ] = v
	return v, nil
}

type Board struct {
	points map[int]map[int]Point
}

func (b *Board) AddPoint(x, y int) {
	m, ok := b.points[x]
	if !ok || m == nil {
		m = make(map[int]Point)
		b.points[x] = m
	}
	m[y] = Point{
		PointType: randomizePointType(),
		x:         x,
		y:         y,
	}
}

func (b Board) Info() map[string]int {
	count := map[string]int{
		"xred":  0,
		"xblue": 0,
		"ored":  0,
		"oblue": 0,
	}
	for _, m := range b.points {
		for _, p := range m {
			count[p.PointType.GetKey()]++
		}
	}
	return count
}

// main
var (
	r                = rand.New(rand.NewSource(time.Now().UnixMilli()))
	pointTypeFactory = &PointTypeFactory{
		pointTypes: make(map[string]PointType),
	}
)

func randomizePointType() PointType {
	keys := []string{
		"xred", "xblue", "ored", "oblue",
	}
	pointType, _ := pointTypeFactory.getPointType(keys[r.Intn(len(keys))])
	return pointType
}

func main() {
	board := Board{
		points: make(map[int]map[int]Point),
	}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			board.AddPoint(i, j)
		}
	}

	count := board.Info()
	log.Printf("fill full board by: %v", count)
}
