package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	size float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (s *Square) Perimeter() float64 {
	return 4 * s.size
}

func (t *Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (s *Square) Area() float64 {
	return s.size * s.size
}

func (t *Triangle) Area() float64 {
	p := t.Perimeter() / 2

	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func Perimeter(f Figure) float64 {
	return f.Perimeter()
}

func Area(f Figure) float64 {
	return f.Area()
}

func main() {
	s := &Square{size: 2}
	t := &Triangle{a: 2, b: 2, c: 2 * math.Sqrt2}
	fmt.Println("Square Area:", s.Area())
	fmt.Println("Triangle Area:", t.Area())
}
