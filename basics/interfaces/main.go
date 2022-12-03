package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) area() float64 {
	return r.width * r.height
}

func (c *Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type Shape interface {
	area() float64
}

func main() {
	c1 := &Circle{4.5}
	c2 := &Rect{5, 7}
	shapes := []Shape{c1, c2}
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
