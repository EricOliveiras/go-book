package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Whell struct {
	Circle
	Spokes int
}

func main() {
	var roda Whell

	roda.Circle.X = 8
	roda.Circle.Y = 8
	roda.Circle.Radius = 5
	roda.Spokes = 20

	w := Whell{
		Circle: Circle{
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)
}