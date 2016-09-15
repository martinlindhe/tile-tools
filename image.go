package tiletools

import (
	"bufio"
	"image"
	"image/draw"
	"log"
	"os"
)

// Section ...
type Section int

// ...
const (
	Top Section = iota
	Bottom
	Left
	Right
)

// KeepSize ...
type KeepSize int

// ...
const (
	OneThird KeepSize = iota
	Half
	TwoThirds
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

// GetPartOfImage returns 1/3:rd of section
func GetPartOfImage(fileName string, section Section, keepSize KeepSize) *image.RGBA {

	img, _, err := DecodeImage(fileName)
	if err != nil {
		panic(err)
	}
	b := img.Bounds()

	var tX, tY int
	var sr image.Rectangle

	if section == Bottom || section == Top {
		switch keepSize {
		case OneThird:
			tX = b.Max.X
			tY = b.Max.Y / 3
		case Half:
			tX = b.Max.X
			tY = b.Max.Y / 2
		case TwoThirds:
			tX = b.Max.X
			tY = (b.Max.Y / 3) * 2
		}
	} else {
		switch keepSize {
		case OneThird:
			tX = b.Max.X / 3
			tY = b.Max.Y
		case Half:
			tX = b.Max.X / 2
			tY = b.Max.Y
		case TwoThirds:
			tX = (b.Max.X / 3) * 2
			tY = b.Max.Y
		}
	}

	switch section {
	case Bottom:
		switch keepSize {
		case OneThird:
			sr = image.Rect(0, tY*2, b.Max.X, b.Max.Y)
		case Half:
			sr = image.Rect(0, b.Max.Y-tY, b.Max.X, b.Max.Y)
		case TwoThirds:
			sr = image.Rect(0, b.Max.Y-tY, b.Max.X, b.Max.Y)
		}

	case Top:
		switch keepSize {
		case OneThird:
			sr = image.Rect(0, 0, b.Max.X, tY)
		case Half:
			sr = image.Rect(0, 0, b.Max.X, tY)
		case TwoThirds:
			sr = image.Rect(0, 0, b.Max.X, tY)
		}

	case Left:
		switch keepSize {
		case OneThird:
			sr = image.Rect(0, 0, b.Max.X-(tX*2), b.Max.Y)
		case Half:
			sr = image.Rect(0, 0, b.Max.X-tX, b.Max.Y)
		case TwoThirds:
			sr = image.Rect(0, 0, tX, b.Max.Y)
		}

	case Right:
		switch keepSize {
		case OneThird:
			sr = image.Rect(b.Max.X-tX, 0, b.Max.X, b.Max.Y)
		case Half:
			sr = image.Rect(tX, 0, b.Max.X, b.Max.Y)
		case TwoThirds:
			sr = image.Rect(b.Max.X-tX, 0, b.Max.X, b.Max.Y)
		}

	default:
		log.Fatal("not handled", section)
		return nil
	}

	dst := image.NewRGBA(image.Rect(0, 0, tX, tY))
	r := sr.Sub(sr.Min).Add(image.Point{0, 0})
	draw.Draw(dst, r, img, sr.Min, draw.Src)
	return dst

}
