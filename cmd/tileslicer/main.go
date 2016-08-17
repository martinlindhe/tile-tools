package main

import (
	"fmt"
	"os"

	tiletools "github.com/martinlindhe/tile-tools/lib"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file       = kingpin.Arg("file", "Input png tileset").Required().File()
	outDir     = kingpin.Arg("outdir", "Output dir").Required().String()
	tileWidth  = kingpin.Arg("width", "Tile width").Required().Int()
	tileHeight = kingpin.Arg("height", "Tile height").Required().Int()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
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
