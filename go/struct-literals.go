package main

import "fmt"

type Point struct{ X, Y int }

type Circle struct {
	*Point
	Radius int
}

func main() {
	p1 := Point{1, 2}

	c1 := Circle{&p1, 44}

	fmt.Println(p1, c1)
	fmt.Println(c1.Point.X)
	fmt.Println(c1.X)

	p2 := new(Point)
	*p2 = Point{1, 2}

	c2 := Circle{p2, 62}

	fmt.Println(c2)

	p3 := &Point{4, 4}

	c4 := Circle{p3, 12}

	fmt.Println(c4)
}

// type Point struct{ X, Y int }

// type Wheel struct {
// 	Circle
// 	//Circle Circle
// 	Spokes int
// }

// func main() {
// 	p1 := Point{1, 2}
// 	p2 := Point{X: 1, Y: 2}

// 	pp := new(Point)
// 	*pp = Point{1, 2}

// 	pp2 := &Point{1, 2}

// 	var w Wheel

// 	w.Circle.Point.X = 8
// 	w.X = 8
// 	w.Circle.Point.Y = 8
// 	w.Y = 8
// 	w.Circle.Radius = 5
// 	w.Radius = 5
// 	w.Spokes = 20

// 	p := new(Point)
// 	*p = Point{8, 8}

// 	//	w3 := Wheel{Circle{p, 5}, 20}

// 	w4 := Wheel{
// 		Circle: Circle{
// 			Point:  p,
// 			Radius: 5,
// 		},
// 		Spokes: 20,
// 	}

// 	fmt.Println("test")
// 	fmt.Println(p1, p2, pp2, w4)

// 	//fmt.Println(p1, p2, pp, pp2, w, w3, w4)
// }
