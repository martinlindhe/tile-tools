package main

import (
	"fmt"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inFile     = kingpin.Arg("in", "Input PNG.").Required().String()
	outFile    = kingpin.Flag("out", "Output PNG.").Required().Short('o').String()
	vertical   = kingpin.Flag("vertical", "Vertical flip.").Short('v').Bool()
	horizontal = kingpin.Flag("horizontal", "Horizontal flip.").Short('h').Bool()
)

func main() {

	kingpin.Parse()

	if !*vertical && !*horizontal {
		fmt.Println("error: either --horizontal or --vertical is required")
		os.Exit(1)
	}

	img, err := imaging.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}

	if *horizontal {
		img = imaging.FlipH(img)
	}

	if *vertical {
		img = imaging.FlipV(img)
	}

	if err := imaging.Save(img, *outFile); err != nil {
		log.Fatal(err)
	}
}
