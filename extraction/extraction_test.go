package extraction

import (
	"reflect"
	"testing"
)

func TestNewExtractionService(t *testing.T) {
	testCases := []struct {
		name string
		f    FFMPEGExtractor
		want ExtractionService
	}{
		{
			name: "successful-service-creation",
			f: FFMPEGExtractor{
				Path: "ffmpeg",
			},
			want: ExtractionService{
				extractor: FFMPEGExtractor{
					Path: "ffmpeg",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			es := NewExtractionService(tc.f)

			if !reflect.DeepEqual(tc.want, es) {
				t.Errorf("got '%v', want '%v", es, tc.want)
			}
		})
	}
}

func TestIFrames(t *testing.T) {
	testCases := []struct {
		name  string
		f     FFMPEGExtractor
		n     int
		video string
		dir   string
	}{
		{
			name: "successful-iframes",
			f: FFMPEGExtractor{
				Path: "ffmpeg",
			},
			n:     10,
			video: "../data/yosemiteA.mp4",
			dir:   "../data/iframes/",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			es := NewExtractionService(tc.f)

			err := es.IFrames(tc.n, tc.video, tc.dir)

			if err != nil {
				t.Errorf("error in iFrames: err = %v", err)
			}
		})
	}
}

func TestGenerateFrames(t *testing.T) {
	testCases := []struct {
		name  string
		f     FFMPEGExtractor
		n     int
		video string
		dir   string
	}{
		{
			name: "successful-generate-frames",
			f: FFMPEGExtractor{
				Path: "ffmpeg",
			},
			n:     10,
			video: "../data/yosemiteA.mp4",
			dir:   "../data/iframes/",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.f.GenerateFrames(tc.n, tc.video, tc.dir)

			if err != nil {
				t.Errorf("error in Generateframes: err = %v", err)
			}
		})
	}
}
