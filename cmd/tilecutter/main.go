package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/martinlindhe/tile-tools"
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

		if err := imaging.Save(img, p); err != nil {
			log.Fatal(err)
		}
	}
}
