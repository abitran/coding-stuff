package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	data := make([][]uint8, dy)
	for x := 0; x < dx; x++ {
		data[x] = make([]uint8, dx)
		for y := 0; y < dy; y++ {
			data[x][y] = uint8((x * y) ^ 3)
		}
	}

	return data
}

func main() {
	pic.Show(Pic)
}
