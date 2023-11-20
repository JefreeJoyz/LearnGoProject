package main

import (
	"fmt"
	"math"
)

type Shape interface { // Это "протокол из свифта"
	Area() float32
}

type Square struct {
	side float32
}

func (s Square) Area() float32 {
	return s.side * s.side
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func testInterface() {
	square := Square{2}
	circle := Circle{3}

	printShapeArea(square)
	printShapeArea(circle)
}

func printShapeArea(shape Shape) { // реализуем интерфейс
	fmt.Println(shape.Area())
}

// Shape2 - интерфейс может включать в себя другие интерфейсы, например:
type Shape2 interface {
	Area() float32
}

type Shape3 interface {
	Shape2
	Shape
}
