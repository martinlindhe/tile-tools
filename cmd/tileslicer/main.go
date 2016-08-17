package main

import (
	"fmt"
	"os"

	tiletools "github.com/martinlindhe/tile-tools/lib"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file       = kingpin.Arg("in", "Input PNG").Required().File()
	outDir     = kingpin.Flag("out", "Output directory").Required().Short('o').String()
	tileWidth  = kingpin.Flag("width", "Tile width").Required().Short('w').Int()
	tileHeight = kingpin.Flag("height", "Tile height").Required().Short('h').Int()
)

func main() {

	kingpin.Parse()

	inFileName := (*file).Name()

	if tiletools.PathDontExist(*outDir) {
		err := os.Mkdir(*outDir, 0777)
		if err != nil {
			fmt.Printf("Could not create %s: %s", *outDir, err)
			os.Exit(1)
		}
	}

	tiletools.SliceImage(inFileName, *outDir, *tileWidth, *tileHeight)
}
