package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

type RGBAPoint struct {
	Point
	color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{&Point{1, 1}, red}
	var q = ColoredPoint{&Point{5, 4}, blue}

	fmt.Println(p.Distance(*q.Point))

	rgbap := RGBAPoint{Point: Point{1, 2}, RGBA: color.RGBA{0, 0, 0, 0}}

	fmt.Println(rgbap)
	fmt.Println(rgbap.RGBA)
}
