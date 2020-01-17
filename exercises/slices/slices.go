package slices

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	addSlicesOfLengthTo(image, dx)
	return image
}

func addSlicesOfLengthTo(image [][]uint8, dx int) {
	for x := range image {
		nestedSlice := make([]uint8, dx)
		for y := range nestedSlice {
			nestedSlice[y] = uint8((x ^ y) * (x ^ y))
		}
		image[x] = nestedSlice
	}
}

func main() {
	pic.Show(Pic)
}
