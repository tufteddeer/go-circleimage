package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func loadPng(name string) (image.Image, error) {
	sourceImageFile, err := os.Open(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sourceImageFile.Close()

	return png.Decode(sourceImageFile)

}

func main() {

	source, err := loadPng("gopher.png")
	if err != nil {
		log.Fatal(err.Error())
	}

	radius := 150
	rounded := CircleImage(source, image.Point{150, 150}, radius)

	outputFile, err := os.Create("result.png")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer outputFile.Close()

	png.Encode(outputFile, rounded)
}
