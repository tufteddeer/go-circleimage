package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

func loadPng(name string) (image.Image, error) {
	sourceImageFile, err := os.Open(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sourceImageFile.Close()

	return png.Decode(sourceImageFile)

}
func main() {

	source, err := loadPng("test.png")
	if err != nil {
		log.Fatal(err.Error())
	}
	
	out := image.NewRGBA(source.Bounds())
	draw.DrawMask(out, source.Bounds(), source, image.ZP, &circle{image.Point{source.Bounds().Dx() / 2, source.Bounds().Dx() / 2}, 100}, image.ZP, draw.Over)

	outputFile, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	png.Encode(outputFile, out)
}
