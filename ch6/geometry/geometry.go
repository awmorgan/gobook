package geometry

import (
	"math"
)

// Point is a set of x, y coordinates
type Point struct{ X, Y float64 }

// Distance traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance same thing, but as a method of a Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// PathDistance returns the distance along the path
func PathDistance(path Path) float64 {
	return path.Distance()
}

// ScaleBy scales a point by a factor
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}