package main

import (
	"os"
	"fmt"
	"log"
	"image/jpeg"
	"github.com/nfnt/resize"
	"image"
	"math"
)

func main() {
	args := os.Args[1:]
	refImageFilename := args[0]
	compImageFilename := args[1]

	normalizedWidth := uint(800)
	normalizedHeight := uint(800)

	referenceImage := loadImage(refImageFilename)
	compareToImage := loadImage(compImageFilename)

	normReferenceImage := resize.Resize(normalizedWidth, normalizedHeight, referenceImage, resize.Lanczos3)
	normCompareToImage := resize.Resize(normalizedWidth, normalizedHeight, compareToImage, resize.Lanczos3)

	maxUInt16 := 65535

	totalDiff := 0

	for x := 0; x < int(normalizedWidth); x++ {
		for y := 0; y < int(normalizedHeight); y++ {
			refR, refG, refB, _ := normReferenceImage.At(x, y).RGBA()
			compR, compG, compB, _ := normCompareToImage.At(x, y).RGBA()

			normRefR := float32(refR)/float32(maxUInt16)
			normRefG := float32(refG)/float32(maxUInt16)
			normRefB := float32(refB)/float32(maxUInt16)
			compRefR := float32(compR)/float32(maxUInt16)
			compRefG := float32(compG)/float32(maxUInt16)
			compRefB := float32(compB)/float32(maxUInt16)

			diff := (normRefR - compRefR) + (normRefG - compRefG) + (normRefB - compRefB)
			totalDiff += int(diff)
		}
	}

	if int(math.Abs(float64(totalDiff))) < 100 {
		fmt.Print("The two images are the same!")
	} else {
		fmt.Print("The two images are not the same!")
	}
}

func loadImage(filename string) image.Image {
	reader, err := os.Open(filename)

	if(err != nil) {
		log.Fatal(err)
	}
	defer reader.Close()

	image, err := jpeg.Decode(reader)

	if(err != nil) {
		log.Fatal(err)
	}

	return image
}