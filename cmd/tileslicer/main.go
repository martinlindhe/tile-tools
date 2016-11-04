package main

import (
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
	tiletools.SliceImage(inFileName, *outDir, *tileWidth, *tileHeight, *force)
}
