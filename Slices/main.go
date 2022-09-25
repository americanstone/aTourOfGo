package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dx)
	for x := range r {
		r[x] = make([]uint8, dy)
		for y := range r[x] {
			r[x][y] = uint8(x ^ y)
		}
	}
	return r
}

func main() {
	pic.Show(Pic)
}
