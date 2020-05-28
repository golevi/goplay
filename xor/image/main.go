package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"os"
	"time"
)

// XorBy is what we used to xor each color of each pixel
type XorBy struct {
	R, G, B uint32
}

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

	var xorbys = []XorBy{}

	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			r, g, b, a := logo.At(w, h).RGBA()

			rxor := rnd.Uint32()
			gxor := rnd.Uint32()
			bxor := rnd.Uint32()

			r = r ^ rxor
			g = g ^ gxor
			b = b ^ bxor

			// fmt.Println(rxor)

			xorby := XorBy{
				R: rxor,
				G: gxor,
				B: bxor,
			}

			xorbys = append(xorbys, xorby)
			// xorbys[width][height] = xorby

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
	fmt.Println(xorbys)
}
