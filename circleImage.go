package main

import (
	"image"
	"image/color"
	"image/draw"
)

/*
 * Based on this article: https://blog.golang.org/go-imagedraw-package
 */

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

// CircleImage creates a new image from a round shape within the original
func CircleImage(source image.Image, origin image.Point, r int) image.Image {
	c := &circle{origin, r}
	result := image.NewRGBA(c.Bounds())

	draw.DrawMask(result, source.Bounds(), source, image.ZP, c, image.ZP, draw.Over)

	return result
}
