package main

import (
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

type Image struct{
    w, h int
}

func (Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
    v := uint8(x+y)
    return color.RGBA{v, v, 255, 255}
}

func main() {
    m := Image{128, 128}
    pic.ShowImage(m)
}
