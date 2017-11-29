package main

import (
	"fmt"
	"log"

	"github.com/martinlindhe/tile-tools"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file       = kingpin.Arg("in", "Input PNG.").Required().File()
	outDir     = kingpin.Flag("out", "Output directory.").Required().Short('o').String()
	tileWidth  = kingpin.Flag("width", "Tile width.").Required().Short('w').Int()
	tileHeight = kingpin.Flag("height", "Tile height.").Required().Short('h').Int()
	force      = kingpin.Flag("force", "Force operation").Bool()
)

func main() {
	kingpin.Parse()
	inFileName := (*file).Name()
	images, err := tiletools.SliceImage(inFileName, *tileWidth, *tileHeight, *force)
	if err != nil {
		log.Fatal(err)
	}
	tiletools.WriteImages(images, *outDir)
	fmt.Println(len(images), "tiles written to", *outDir)
}
