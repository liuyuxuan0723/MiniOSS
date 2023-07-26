package demo

import (
	"image"
	"image/color"
)

type Image struct {
	Height, Width int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, uint8(img.Height), uint8(img.Width)}
}
