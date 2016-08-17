# About

Command line tools to process tile set images


### Installation

    go get -u github.com/martinlindhe/tile-tools/...


### Usage - tileslicer

Slice input PNG into several tiles, save to out_font

    $ tileslicer _samples/ground.png --out _tmp --width 8 --height 12
    74 tiles written to _tmp


### Usage - tilecutter

Keeps bottom 1/3:rd of input image

    $ tilecutter _tmp

TODO: allow to choose bottom or top part, allow to choose half, 1/3 or 2/3 to keep


### Usage - tilejoiner

Combines all PNG:s in a folder into a new tile set

    $ tilejoiner _tmp --out ground-out.png --tiles-per-row 4
    Creating tileset of 74 tiles with 8x4 pixels, 4 tiles per row. Output is image is 32x32 pixels
    Writing to ground-out.png


### License

Under [MIT](LICENSE)
