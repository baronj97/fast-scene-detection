// Package frames contains an API for interacting with i-frames and representing them into some data structure.
package frames

import (
	"image"
	"io"

	// Due to the way the image package works, we need to register the png import as well, even if it isn't used directly
	_ "image/png"
)

// CompressionType represents the type of Compression used on the video file.
type CompressionType int

const (
	// MPEG2 is used for files saved with .mp2 and have i-frames every 0.5 seconds.
	MPEG2 CompressionType = iota
	// MPEG4 is used for files saved with .mp4 and have i-frames every 0.20002 seconds.
	MPEG4
)

// Processor is an interface for processing i-frames into exported Frame structures. The interface requires a single
// method, Process, which accepts a filename and returns a Frame type to represent it.
type Processor interface {
	Process(io.Reader) (uint32, uint32, uint32, uint32, error)
}

// Frame represents an i-frame by maintaining a single value for each color channel and the number of frame.
type Frame struct {
	num int    // The frame number within the video
	r   uint32 // Red color channel value
	g   uint32 // Green color channel value
	b   uint32 // Blue color channel value
	a   uint32 // Alpha blending factor value
}

// NewFrame creates a Frame structure.
func NewFrame(num int, r, g, b, a uint32) Frame {
	return Frame{
		num: num,
		r:   r,
		g:   g,
		b:   b,
		a:   a,
	}
}

// Red returns the value of the red color channel.
func (f Frame) Red() uint32 {
	return f.r
}

// Green returns the value of the green color channel.
func (f Frame) Green() uint32 {
	return f.g
}

// Blue returns the value of the blue color channel.
func (f Frame) Blue() uint32 {
	return f.b
}

// Num returns the number of the frame.
func (f Frame) Num() int {
	return f.num
}

// Alpha returns the value of the alpha blending factor.
func (f Frame) Alpha() uint32 {
	return f.a
}

// FrameService exports the interface through a pluggable interface abstraction.
type FrameService struct {
	processor Processor
}

// NewFrameService returns a FrameService composed of a processor p.
func NewFrameService(p Processor) FrameService {
	return FrameService{
		processor: p,
	}
}

// Read accepts the path to a file and calls the processor to process the path.
func (fs FrameService) Read(r io.Reader) (uint32, uint32, uint32, uint32, error) {
	return fs.processor.Process(r)
}

// FrameReader is an exported structure to work on the i-frames.
type FrameReader struct {
	Compression CompressionType
	Compute     func(r io.Reader) (uint32, uint32, uint32, uint32, error)
}

// Process processess the filepath and returns the Frame.
func (fr FrameReader) Process(r io.Reader) (uint32, uint32, uint32, uint32, error) {
	return fr.Compute(r)
}

// Average is an exported processing function which can be attahced to the FrameReader.
func Average(r io.Reader) (uint32, uint32, uint32, uint32, error) {
	m, _, err := image.Decode(r)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	dims := m.Bounds().Max
	var rBucket, gBucket, bBucket, aBucket uint32
	for i := 0; i < dims.X; i++ {
		for j := 0; j < dims.Y; j++ {
			rm, gm, bm, am := m.At(i, j).RGBA()
			rBucket += rm
			bBucket += bm
			gBucket += gm
			aBucket += am
		}
	}
	return rBucket / uint32(dims.X*dims.Y), gBucket / uint32(dims.X*dims.Y), bBucket / uint32(dims.X*dims.Y), aBucket / uint32(dims.X*dims.Y), nil
}
