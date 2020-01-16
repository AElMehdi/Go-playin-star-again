package image

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	x, y        int
	blue, alpha uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x, i.y)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), i.blue, i.alpha}
}

func main() {
	m := Image{500, 300, 222, 222}
	pic.ShowImage(m)
}
