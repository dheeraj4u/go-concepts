package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mat := make([][]uint8, dy)
	for i := range mat {
		mat[i] = make([]uint8, dx)
		for j := range mat[i] {
			mat[i][j] = uint8((i + j) / 2) //x*y //x^y
		}
	}
	return mat
}

func main() {
	pic.Show(Pic)
}
