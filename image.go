package tiletools

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/disintegration/imaging"
)

// Section indicates which part of the image to keep
type Section int

// ...
const (
	Top Section = iota
	Bottom
	Left
	Right
)

// KeepSize indicates the size of the image to keep
type KeepSize int

// ...
const (
	OneThird KeepSize = iota
	Half
	TwoThirds
)

// LoadImage loads a image from disk
func LoadImage(filename string) (image.Image, string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return image.Decode(bufio.NewReader(f))
}

// WriteImages saves a slice of images to disk
func WriteImages(imgs []image.Image, dstDir string) error {
	mkdirIfNotExisting(dstDir)
	cnt := 0
	for _, img := range imgs {
		if rgba, ok := img.(*image.RGBA); ok {
			if isOnlyTransparent(rgba) {
				continue
			}
		}
		fileName := fmt.Sprintf("%s/%03d.png", dstDir, cnt)
		if err := imaging.Save(img, fileName); err != nil {
			return err
		}
		cnt++
	}
	return nil
}

func usedBounds(section Section, keepSize KeepSize, b image.Rectangle) (int, int) {
	var tX, tY int
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
	return tX, tY
}

func makeRect(section Section, keepSize KeepSize, b image.Rectangle) (int, int, image.Rectangle) {
	var sr image.Rectangle
	tX, tY := usedBounds(section, keepSize, b)

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
	}
	return tX, tY, sr
}

// GetPartOfImage returns specified Section and Size of image
func GetPartOfImage(fileName string, section Section, keepSize KeepSize) (*image.RGBA, error) {
	img, _, err := LoadImage(fileName)
	if err != nil {
		return nil, err
	}
	b := img.Bounds()
	tX, tY, sr := makeRect(section, keepSize, b)
	dst := image.NewRGBA(image.Rect(0, 0, tX, tY))
	r := sr.Sub(sr.Min).Add(image.Point{0, 0})
	draw.Draw(dst, r, img, sr.Min, draw.Src)
	return dst, nil
}
