package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y int }

type Path []Point

func (p Point) Distance(q Point) float64 { return math.Hypot(float64(q.X-p.X), float64(q.Y-p.Y)) }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Point{1, 2}
	q := Point{3, 4}

	distanceFromP := p.Distance

	fmt.Println(distanceFromP(q))

	origin := Point{}

	fmt.Println(distanceFromP(origin))

	distance := Point.Distance

	fmt.Println(distance(p, q))

	punkte := make(Path, 5, 10)

	punkte[0] = Point{3, 4}
	punkte[1] = Point{3, 4}
	punkte[2] = Point{3, 4}

	punkte.TranslateBy(punkte[1], true)
}
