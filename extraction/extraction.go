package extraction

import (
	"os/exec"
	"strconv"
)

// Extractor is an abstracted interface for generating the i-frames. It contains a single method,
// GenerateFrames, which accepts the interval of i-frames to take, the path to the video file, and
// the output directory.
type Extractor interface {
	GenerateFrames(n int, video string, dir string) error
}

// ExtractionService hides the Extractor interface.
type ExtractionService struct {
	extractor Extractor
}

// NewExtractionService creates a new ExtractionService structure.
func NewExtractionService(e Extractor) ExtractionService {
	return ExtractionService{
		extractor: e,
	}
}

// ExtractionService exports the GenerateFrames function.
func (ex ExtractionService) IFrames(n int, video, dir string) error {
	return ex.extractor.GenerateFrames(n, video, dir)
}

// FFMPEGExtractor is an exported structure that holds the path to the executable for FFMPEG.
type FFMPEGExtractor struct {
	Path string
}

// GenerateFrames is an unexported processing function attached to FMMPEGExtractor. It will
// utilize FFMPEG in order to extract a frame for every n frames in the video and store the
// ouputs to the dir path.
func (f FFMPEGExtractor) GenerateFrames(n int, video string, dir string) error {
	cmd := exec.Command("ffmpeg", "-i", video, "-vf", "select=not(mod(n\\,"+strconv.Itoa(n)+"))",
		"-vsync", "vfr", dir+"frame%03d.png")

	return cmd.Run()
}
