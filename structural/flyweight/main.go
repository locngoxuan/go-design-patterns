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

var pointTypes = map[string]PointType{
	"xred": {
		Color: "red",
		Char:  "x",
	},
	"xblue": {
		Color: "blue",
		Char:  "x",
	},
	"ored": {
		Color: "red",
		Char:  "o",
	},
	"oblue": {
		Color: "blue",
		Char:  "o",
	},
}

type Board struct {
	points map[int]map[int]Point
}

var r = rand.New(rand.NewSource(time.Now().UnixMilli()))

func randomizePointType() PointType {
	keys := []string{
		"xred", "xblue", "ored", "oblue",
	}
	return pointTypes[keys[r.Intn(len(keys))]]
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
