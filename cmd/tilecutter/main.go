package main

import (
	"io/ioutil"
	"path/filepath"

	tiletools "github.com/martinlindhe/tile-tools/lib"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inDir = kingpin.Arg("in", "Input directory").Required().String()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// loop over input folder, keep bottom third of each image, overwrite
	files, _ := ioutil.ReadDir(*inDir)
	for _, f := range files {

		p := filepath.Join(*inDir, f.Name())
		img := tiletools.GetBottomThirdOfImage(p)

		tiletools.WritePNG(p, img)
	}
}
