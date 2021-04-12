package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mat := make([][]uint8, dy)
	for x := range mat {
		mat[x] = make([]uint8, dx)
		for y := range mat[x] {
			mat[x][y] = uint8(x * y)
		}
	}
	return mat
}

func main() {
	pic.Show(Pic)
}
