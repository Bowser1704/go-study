package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	xx := make([][]uint8, dy)
	for i := range xx {
		xx[i] = make([]uint8, dx)
		for j := range xx[i]{
			xx[i][j] = uint8(i+j)
		}
	}
	return xx
}

func main() {
	pic.Show(Pic)
}
