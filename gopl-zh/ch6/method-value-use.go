package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

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
	path := Path{Point{1, 2}, Point{3, 4}, Point{5, 6}}
	offst := Point{-1, 1}
	path.TranslateBy(offst, true)
	fmt.Println(path)
	path.TranslateBy(offst, false)
	fmt.Println(path)
}
