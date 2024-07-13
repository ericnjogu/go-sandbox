package lib

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.Radius * 2
}

type Rect struct {
	Len   float64
	Width float64
}

func (r Rect) area() float64 {
	return r.Len * r.Width
}

func (r Rect) perimeter() float64 {
	return (r.Width + r.Len) * 2
}

func Measure(g Geometry) {
	fmt.Printf("%T: %+v\n", g, g)
	fmt.Println("area", g.area())
	fmt.Println("perimeter", g.perimeter())
}
