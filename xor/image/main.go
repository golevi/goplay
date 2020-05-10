package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"os"
	"time"
)

func main() {
	logoFile, err := os.Open("logo.jpg")
	if err != nil {
		panic(err)
	}
	defer logoFile.Close()

	logo, err := jpeg.Decode(logoFile)
	if err != nil {
		panic(err)
	}
	width := logo.Bounds().Max.X
	height := logo.Bounds().Max.Y

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			r, g, b, a := logo.At(w, h).RGBA()

			r = r ^ rnd.Uint32()
			g = g ^ rnd.Uint32()
			b = b ^ rnd.Uint32()

			c := color.RGBA{
				R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a),
			}
			dst.Set(w, h, c)
		}
	}

	dstfp, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer dstfp.Close()

	err = jpeg.Encode(dstfp, dst, nil)
	if err != nil {
		panic(err)
	}
}
