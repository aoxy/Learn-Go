package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["USTC"] = Vertex{
		117.26139, 31.83819,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["USTC"])
}
