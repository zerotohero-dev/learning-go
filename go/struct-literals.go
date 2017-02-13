package main

import "fmt"

type Point struct{ X, Y int }

type Circle struct {
	//Center Point
	Point
	Radius int
}

type Wheel struct {
	Circle
	//Circle Circle
	Spokes int
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{X: 1, Y: 2}

	pp := new(Point)
	*pp = Point{1, 2}

	pp2 := &Point{1, 2}

	var w Wheel

	w.Circle.Point.X = 8
	w.X = 8
	w.Circle.Point.Y = 8
	w.Y = 8
	w.Circle.Radius = 5
	w.Radius = 5
	w.Spokes = 20

	w3 := Wheel{Circle{Point{8, 8}, 5}, 20}

	w4 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Println(p1, p2, pp, pp2, w, w3, w4)
}
