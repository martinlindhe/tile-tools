package tiletools

import (
	"fmt"
	"image"
	"image/draw"
	"math"
)

// SliceImage slices input imgFile into smaller tiles
func SliceImage(imgFile string, tileWidth int, tileHeight int, force bool) ([]image.Image, error) {
	var slices []image.Image

	img, _, err := LoadImage(imgFile)
	if err != nil {
		return nil, err
	}

	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	cols := float64(imgWidth) / float64(tileWidth)
	rows := float64(imgHeight) / float64(tileHeight)

	if !force && cols != math.Floor(cols) {
		return nil, fmt.Errorf("input image width %d is not evenly divisible by tile width %d", imgWidth, tileWidth)
	}

	if !force && rows != math.Floor(rows) {
		return nil, fmt.Errorf("input image height %d is not evenly divisible by tile height %d", imgHeight, tileHeight)
	}

	for row := 0; row < int(rows); row++ {
		for col := 0; col < int(cols); col++ {
			x0 := col * tileWidth
			y0 := row * tileHeight
			x1 := (col + 1) * tileWidth
			y1 := (row + 1) * tileHeight
			sr := image.Rect(x0, y0, x1, y1)

			dst := image.NewRGBA(image.Rect(0, 0, tileWidth, tileHeight))
			r := sr.Sub(sr.Min).Add(image.Point{0, 0})
			draw.Draw(dst, r, img, sr.Min, draw.Src)
			slices = append(slices, dst)
		}
	}
	return slices, nil
}

// is this an empty tile?
func isOnlyTransparent(img *image.RGBA) bool {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				return false
			}
		}
	}
	return true
}
