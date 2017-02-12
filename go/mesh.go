package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyRange       = 30.0
	xyScale       = width / 2 / xyRange
	zScale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf(
		"<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' "+
			" width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Printf("<polygon points='%g, %g %g, %g %g, %g %g, %g' />", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)

	return math.Sin(r) / r
}
