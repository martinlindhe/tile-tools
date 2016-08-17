package tiletools

import (
	"bufio"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// DecodeImage ...
func DecodeImage(filename string) (image.Image, string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return image.Decode(bufio.NewReader(f))
}

// WritePNG ...
func WritePNG(fileName string, m image.Image) error {

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}

// GetBottomThirdOfImage ...
func GetBottomThirdOfImage(fileName string) *image.RGBA {

	img, _, err := DecodeImage(fileName)
	if err != nil {
		panic(err)
	}

	b := img.Bounds()

	tX := b.Max.X
	tY := b.Max.Y / 3

	sr := image.Rect(0, tY*2, b.Max.X, b.Max.Y)

	dst := image.NewRGBA(image.Rect(0, 0, tX, tY))
	r := sr.Sub(sr.Min).Add(image.Point{0, 0})
	draw.Draw(dst, r, img, sr.Min, draw.Src)

	return dst
}
