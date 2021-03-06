package main

import (
	"fmt"
	"image/color"
	"math"
)

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = coloredpoint{point{1, 1}, red}
	var q = coloredpoint{point{5, 4}, blue}
	fmt.Println(p.Distance(q.point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.point))
}

type point struct{ X, Y float64 }

type coloredpoint struct {
	point
	Color color.RGBA
}

func (p point) Distance(q point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func init() {
	p := point{1, 2}
	q := point{4, 6}
	distance := point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	type coloredpoint struct {
		*point
		Color color.RGBA
	}

	p := coloredpoint{&point{1,1}, red}
	q := coloredpoint{&point{5,4}, blue}

	fmt.Println(p.Distance(*q.point))
	q.point = p.point
	p.ScaleBy(2)
	fmt.Println(*p.point, *q.point)
}
