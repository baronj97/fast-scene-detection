package main

import (
	"fmt"

	"github.com/baronj97/fast-scene-detection/extraction"
)

func main() {
	// FFMPEGExtractor is an exported structure that holds the path to the exdecutable for FFMPEG.
    e := extraction.FFMPEGExtractor{Path: "ffmpeg"}
	es := extraction.NewExtractionService(e)
	es.IFrames(100, "/usr/local/go/src/fast-scene-detection/data/yosemiteA.mp4", "/usr/local/go/src/fast-scene-detection/data/iframes/")

	//extraction.IFrames(10, "../data/yosemiteA", "../data/iframes")
	fmt.Println("DONE")
}