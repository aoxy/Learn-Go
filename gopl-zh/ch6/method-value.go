package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance

	fmt.Println(distanceFromP(q))

	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	fmt.Println(p)
	scaleP(2)
	fmt.Println(p)
	scaleP(3)
	fmt.Println(p)
	scaleP(10)
	fmt.Println(p)

	distance := Point.Distance
	fmt.Printf("%T\n", distance)
	fmt.Println(distance(p, q))

	scale := (*Point).ScaleBy
	scale(&q, 3)
	fmt.Printf("%T\n", scale)
	fmt.Println(q)
}
