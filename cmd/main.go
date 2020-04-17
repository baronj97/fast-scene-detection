package main

import (
	//"log"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strconv"
	//"bytes"
	//"os/exec"
)

func main() {
	inFile := "../data/yosemiteA.mp4"
	frameDir := "../data/iframes/"
	n := 1

	cmd := exec.Command("ffmpeg", "-i", inFile, "-vf", "select=not(mod(n\\,"+strconv.Itoa(n)+"))",
		"-vsync", "vfr", frameDir+"frame%03d.png")

	var stdBuffer bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
	// Output:
	log.Println(stdBuffer.String())
}
