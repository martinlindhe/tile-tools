test:
	go run cmd/tileslicer/main.go _samples/ground.png _tmp 8 12
	go run cmd/tilecutter/main.go _tmp
	go run cmd/tilejoiner/main.go _tmp ground-out.png 4
	rm -rf _tmp
#	rm ground.png
