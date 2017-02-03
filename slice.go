package tiletools

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"math"

	"github.com/disintegration/imaging"
)

// SliceImage is used by cmd/tileslicer
func SliceImage(imgFile string, outDir string, tileWidth int, tileHeight int, force bool) []image.Image {
	var slices []image.Image
	mkdirIfNotExisting(outDir)

	img, _, err := DecodeImage(imgFile)
	if err != nil {
		panic(err)
	}

	b := img.Bounds()
	imgWidth := b.Max.X
	imgHeight := b.Max.Y

	cols := float64(imgWidth) / float64(tileWidth)
	rows := float64(imgHeight) / float64(tileHeight)

	if !force && cols != math.Floor(cols) {
		log.Fatalf("Input image width %d is not evenly divisable by tile width %d", imgWidth, tileWidth)
	}

	if !force && rows != math.Floor(rows) {
		log.Fatalf("Input image height %d is not evenly divisable by tile height %d", imgHeight, tileHeight)
	}

	// slice up image into tiles
	cnt := 0
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

			if isOnlyTransparent(dst) {
				fmt.Printf("Skipping empty tile at row %d, col %d\n", row, col)
				continue
			}

			outFile := fmt.Sprintf("%s/%03d.png", outDir, cnt)
			if err := imaging.Save(dst, outFile); err != nil {
				log.Fatal(err)
			}
			cnt++
		}
	}

	fmt.Printf("%d tiles written to %s\n", cnt, outDir)
	return slices
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
