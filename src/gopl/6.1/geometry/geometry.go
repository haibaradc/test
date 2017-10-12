package geometry

import (
	"math"
)

type Point struct {
	X float64
	Y float64
	//X,Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) ScaleBy2(f float64) {
	p.X *= f
	p.Y *= f
}

func (p *Point) ScaleBy(f float64) {
	p.X *= f
	p.Y *= f
}

type Path []Point

func (p Path) Distance() float64 {
	var sum float64
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
