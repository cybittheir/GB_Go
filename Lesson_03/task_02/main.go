package main

import (
	"fmt"
	"math"
)

type Shape struct {
	Name string
}

type Circle struct {
	s      Shape
	radius float64
}

type Rectangle struct {
	s      Shape
	width  float64
	height float64
}

func (s Shape) GetName() string {
	return s.Name
}

func (s Shape) Area() float64 {
	return s.Area()
}

func (c Circle) GetName() string {
	return c.s.Name
}

func (r Rectangle) GetName() string {
	return r.s.Name
}

func (c Circle) Area() float64 {
	return math.Pi * math.Exp2(c.radius)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func main() {
	circle := Circle{Shape{"Круг"}, 5}
	rectangle := Rectangle{Shape{"Прямоугольник"}, 4, 6}

	fmt.Printf("%s: Площадь = %.2f\n", circle.GetName(), circle.Area())
	fmt.Printf("%s: Площадь = %.2f\n", rectangle.GetName(), rectangle.Area())
}
