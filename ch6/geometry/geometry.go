package geometry

import "math"

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
