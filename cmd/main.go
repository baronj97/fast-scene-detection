package main
import (
        "fmt"
        "log"
        "os"
        "bytes"
        "os/exec"
		
        "fast-scene-detection/extraction"
)

func main() {
    // To do:
    a := extraction.Show()
    fmt.Println(a)

    fmt.Println("Running FFMPEG") 
    inFile := "yosemiteA.mp4"
    cmd := exec.Command("ffmpeg", "-i", inFile, "-f", "image2", "-vf", "select='eq(pict_type,PICT_TYPE_I)'", 
                        "-vsync", "vfr", "yi%03d.png")

    var stdBuffer bytes.Buffer
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        log.Panic(err)
    }
    // Output: 
    log.Println(stdBuffer.String())
}
