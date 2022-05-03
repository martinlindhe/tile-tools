test:
	go run cmd/tileslicer/main.go _samples/ground.png --out _tmp --width 8 --height 12
	go run cmd/tilecutter/main.go _tmp --keep-bottom --one-third
	go run cmd/tilejoiner/main.go _tmp --out ground-out.png --tiles-per-row 4
	rm -rf _tmp
#	rm ground-out.png

install:
	go get ./cmd/...
