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
	inDir      = kingpin.Arg("in", "Input directory.").Required().String()
	keepTop    = kingpin.Flag("keep-top", "Keep top part of image.").Bool()
	keepBottom = kingpin.Flag("keep-bottom", "Keep bottom part of image.").Bool()
	keepLeft   = kingpin.Flag("keep-left", "Keep left part of image.").Bool()
	keepRight  = kingpin.Flag("keep-right", "Keep right part of image.").Bool()
	half       = kingpin.Flag("half", "Use 1/2 of image.").Bool()
	oneThird   = kingpin.Flag("one-third", "Use 1/3 of image.").Bool()
	twoThirds  = kingpin.Flag("two-thirds", "Use 2/3 of image.").Bool()
)

func main() {
	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()
	section, keepSize := parseSectionAndKeepSize()

	files, err := ioutil.ReadDir(*inDir)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// loop over input folder
	for _, f := range files {
		p := filepath.Join(*inDir, f.Name())
		img, err := tiletools.GetPartOfImage(p, section, keepSize)
		if err != nil {
			log.Fatal(err)
		}
		if err := imaging.Save(img, p); err != nil {
			log.Fatal(err)
		}
	}
}

func parseSectionAndKeepSize() (tiletools.Section, tiletools.KeepSize) {
	var section tiletools.Section
	var keepSize tiletools.KeepSize
	keep := 0
	size := 0
	if *keepTop {
		section = tiletools.Top
		keep++
	}
	if *keepBottom {
		section = tiletools.Bottom
		keep++
	}
	if *keepRight {
		section = tiletools.Right
		keep++
	}
	if *keepLeft {
		section = tiletools.Left
		keep++
	}
	if *half {
		keepSize = tiletools.Half
		size++
	}
	if *oneThird {
		keepSize = tiletools.OneThird
		size++
	}
	if *twoThirds {
		keepSize = tiletools.TwoThirds
		size++
	}
	checkErrors(keep, size)
	return section, keepSize
}

func checkErrors(keep, size int) {
	switch {
	case keep > 1:
		log.Fatalf("error: use only one of --keep-top, --keep-bottom, --keep-left and --keep-right")
	case keep == 0:
		log.Fatalf("error: use either --keep-top, --keep-bottom, --keep-left or --keep-right")
	case size > 1:
		log.Fatal("error: use only one of --half, --one-third and --two-third")
	case size == 0:
		log.Fatal("error: use either --half, --one-third or --two-thirds")
	}
}
