package main

import (
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"path/filepath"

	tiletools "github.com/martinlindhe/tile-tools/lib"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inDir       = kingpin.Arg("in", "Input directory").Required().String()
	outFile     = kingpin.Flag("out", "Output PNG").Required().Short('o').String()
	tilesPerRow = kingpin.Flag("tiles-per-row", "Tiles per row").Required().Int()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	var images []image.Image

	tileWidth := 0
	tileHeight := 0
	tileCount := 0

	files, _ := ioutil.ReadDir(*inDir)
	for _, f := range files {

		p := filepath.Join(*inDir, f.Name())

		img, _, err := tiletools.DecodeImage(p)
		if err != nil {
			fmt.Printf("Error decoding: %s", err)
			continue
		}

		b := img.Bounds()

		if tileWidth == 0 && tileHeight == 0 {
			tileWidth = b.Max.X
			tileHeight = b.Max.Y
		} else if b.Max.X != tileWidth || b.Max.Y != tileHeight {
			fmt.Printf("Error: tile %s did not have expected dimensions of %d,%d\n", p, tileWidth, tileHeight)
		}

		tileCount++
		images = append(images, img)
	}

	outWidth := *tilesPerRow * tileWidth
	outHeight := (tileCount / *tilesPerRow) * tileHeight

	fmt.Printf("Creating tileset of %d tiles with %dx%d pixels, %d tiles per row. Output is image is %dx%d pixels\n", tileCount, tileWidth, tileHeight, *tilesPerRow, outWidth, outHeight)

	dst := image.NewRGBA(image.Rect(0, 0, outWidth, outHeight))
	for i, img := range images {
		x0 := (i % *tilesPerRow) * tileWidth
		y0 := (i / *tilesPerRow) * tileHeight

		// fmt.Printf("Writing %d to %d,%d\n", i, x0, y0)

		dr := image.Rect(x0, y0, x0+tileWidth, y0+tileHeight)

		draw.Draw(dst, dr, img, image.Point{0, 0}, draw.Src)
	}

	fmt.Printf("Writing to %s\n", *outFile)
	tiletools.WritePNG(*outFile, dst)
}
