package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)
	p := &Vertex{5, 12}
	p.Scale(3)      // p是指针也可以
	ScaleFunc(p, 8) //参数必须是指针
	fmt.Println(v, p)
}
