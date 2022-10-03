package main

import (
	"fmt"
	"log"
	"strings"
)

type Shape interface {
	Move(int, int)
	Draw()
	Accept(Visitor) string
}

type EmptyShape struct {
	Id int
	X  int
	Y  int
}

func (d EmptyShape) Move(int, int) {
	//move
}

func (d EmptyShape) Draw() {
	//draw
}

func (d EmptyShape) Accept(v Visitor) string {
	return ""
}

type Dot struct {
	EmptyShape
}

func (d Dot) Accept(v Visitor) string {
	return v.VisitDot(d)
}

type Circle struct {
	EmptyShape
	Radius int
}

func (d Circle) Accept(v Visitor) string {
	return v.VisitCircle(d)
}

func (d Circle) GetRadius() int {
	return d.Radius
}

type Rectangle struct {
	EmptyShape
	Width  int
	Height int
}

func (d Rectangle) Accept(v Visitor) string {
	return v.VisitRectangle(d)
}

func (d Rectangle) GetWidth() int {
	return d.Width
}
func (d Rectangle) GetHeight() int {
	return d.Height
}

type CompoundShape struct {
	EmptyShape
	Shapes []Shape
}

func (d CompoundShape) Accept(v Visitor) string {
	return v.VisitCompoundGraphic(d)
}

func (d *CompoundShape) Add(s Shape) {
	d.Shapes = append(d.Shapes, s)
}

type Visitor interface {
	VisitDot(Dot) string
	VisitCircle(Circle) string
	VisitRectangle(Rectangle) string
	VisitCompoundGraphic(CompoundShape) string
}

type XMLExportVisitor struct {
}

func (x XMLExportVisitor) Export(shapes ...Shape) string {
	var builder strings.Builder
	builder.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\"?>")
	builder.WriteString("\n")
	for _, shape := range shapes {
		builder.WriteString(shape.Accept(x))
		builder.WriteString("\n")
	}
	return builder.String()
}

func (x XMLExportVisitor) VisitDot(dot Dot) string {
	return fmt.Sprintf(`<dot>
	<id>%v</id>
	<x>%v</x>
	<y>%v</y>
</dot>`, dot.Id, dot.X, dot.Y)
}

func (x XMLExportVisitor) VisitCircle(c Circle) string {
	return fmt.Sprintf(`<circle>
	<id>%v</id>
	<x>%v</x>
	<y>%v</y>
	<radius>%v</radius>
</circle>`, c.Id, c.X, c.Y, c.Radius)
}

func (x XMLExportVisitor) VisitRectangle(r Rectangle) string {
	return fmt.Sprintf(`<rectangle>
	<id>%v</id>
	<x>%v</x>
	<y>%v</y>
	<width>%v</width>
	<height>%v</height>
</rectangle>`, r.Id, r.X, r.Y, r.Width, r.Height)
}

func (x XMLExportVisitor) VisitCompoundGraphic(c CompoundShape) string {
	var builder strings.Builder
	builder.WriteString("<compound_graphic>")
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("	<id>%v</id>", c.Id))
	builder.WriteString("\n")
	builder.WriteString("	<content>")
	builder.WriteString("\n")
	for _, s := range c.Shapes {
		v := s.Accept(x)
		v = strings.ReplaceAll(v, "\n", "\n		")
		v = "		" + v
		builder.WriteString(v)
		builder.WriteString("\n")
	}
	builder.WriteString("	</content>")
	builder.WriteString("\n")
	builder.WriteString("</compound_graphic>")
	return builder.String()
}

func main() {
	dot := Dot{
		EmptyShape: EmptyShape{
			Id: 1,
			X:  10,
			Y:  5,
		},
	}
	circle := Circle{
		EmptyShape: EmptyShape{
			Id: 2,
			X:  23,
			Y:  15,
		},
		Radius: 5,
	}
	rectangle := Rectangle{
		EmptyShape: EmptyShape{
			Id: 3,
			X:  10,
			Y:  17,
		},
		Width:  20,
		Height: 30,
	}
	compoundShape := CompoundShape{
		Shapes: make([]Shape, 0),
	}
	compoundShape.Add(dot)
	compoundShape.Add(circle)
	compoundShape.Add(rectangle)

	xmlExporter := XMLExportVisitor{}
	log.Println(xmlExporter.Export(dot, rectangle, compoundShape, circle))
}
