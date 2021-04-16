package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Splkes int
}

func main() {
	var w Wheel
	w.X = 5
	w.Y = 7
	w.Radius = 9
	w.Splkes = 3

	w2 := Wheel{Circle{Point{5, 7}, 9}, 3}

	w3 := Wheel{
		Circle: Circle{
			Point:  Point{X: 5, Y: 7},
			Radius: 9,
		},
		Splkes: 3,
	}

	fmt.Println(w)
	fmt.Println(w2)
	fmt.Printf("%#v\n", w3)
}
