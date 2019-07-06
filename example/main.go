package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"github.com/tufteddeer/go-circleImage"
)

func loadPng(name string) (image.Image, error) {
	sourceImageFile, err := os.Open(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sourceImageFile.Close()

	return png.Decode(sourceImageFile)

}g

func main() {

	source, err := loadPng("biplane.png")
	if err != nil {
		log.Fatal(err.Error())
	}

	radius := 70
	rounded := circleimage.CircleImage(source, image.Point{115, 80}, radius)

	outputFile, err := os.Create("result.png")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	png.Encode(outputFile, rounded)
}
