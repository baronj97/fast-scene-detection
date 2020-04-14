package frames

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestNewFrame(t *testing.T) {
	testCases := []struct {
		name  string
		id    int
		red   uint32
		green uint32
		blue  uint32
		alpha uint32
		want  Frame
	}{
		{
			name:  "successful-frame",
			id:    1,
			red:   100,
			green: 200,
			blue:  0,
			alpha: 0,
			want: Frame{
				id: 1,
				r:  100,
				g:  200,
				b:  0,
				a:  0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewFrame(tc.id, tc.red, tc.green, tc.blue, tc.alpha)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("got '%v', want '%v", got, tc.want)
			}
		})
	}
}

func TestRed(t *testing.T) {
	testCases := []struct {
		name string
		got  Frame
		want uint32
	}{
		{
			name: "successful-red",
			got: Frame{
				r: 100,
			},
			want: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got.Red() != tc.want {
				t.Errorf("got '%v' want '%v'", tc.got.Red(), tc.want)
			}
		})
	}
}

func TestGreen(t *testing.T) {
	testCases := []struct {
		name string
		got  Frame
		want uint32
	}{
		{
			name: "successful-green",
			got: Frame{
				g: 150,
			},
			want: 150,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got.Green() != tc.want {
				t.Errorf("got '%v' want '%v'", tc.got.Green(), tc.want)
			}
		})
	}
}

func TestBlue(t *testing.T) {
	testCases := []struct {
		name string
		got  Frame
		want uint32
	}{
		{
			name: "successful-blue",
			got: Frame{
				b: 200,
			},
			want: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got.Blue() != tc.want {
				t.Errorf("got '%v' want '%v'", tc.got.Blue(), tc.want)
			}
		})
	}
}

func TestAlpha(t *testing.T) {
	testCases := []struct {
		name string
		got  Frame
		want uint32
	}{
		{
			name: "successful-alpha",
			got: Frame{
				a: 250,
			},
			want: 250,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got.Alpha() != tc.want {
				t.Errorf("got '%v' want '%v'", tc.got.Alpha(), tc.want)
			}
		})
	}
}

func TestID(t *testing.T) {
	testCases := []struct {
		name string
		got  Frame
		want int
	}{
		{
			name: "successful-ID",
			got: Frame{
				id: 1,
			},
			want: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got.ID() != tc.want {
				t.Errorf("got '%v' want '%v'", tc.got.ID(), tc.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	testCases := []struct {
		name     string
		filepath string
		want     Frame
		wantErr  bool
	}{
		{
			name:     "successful-average",
			filepath: "test1.png",
			want: Frame{
				r: 7903,
				g: 7824,
				b: 6843,
				a: 9610,
			},
		},
		{
			name:     "bad-format",
			filepath: "test2.jpg",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := os.Open(tc.filepath)
			if err != nil {
				t.Fatalf("error opening file: err = %v", err)
			}
			defer r.Close()
			gotR, gotG, gotB, gotA, err := Average(r)
			if err != nil {
				if tc.wantErr {
					return
				}
				t.Fatalf("unexpected error: got '%v', want 'nil'", err)
			}
			if gotR != tc.want.Red() {
				t.Errorf("error with red channel: got '%v', want '%v'", gotR, tc.want.Red())
			}

			if gotG != tc.want.Green() {
				t.Errorf("error with green channel: got '%v', want '%v'", gotG, tc.want.Green())
			}

			if gotB != tc.want.Blue() {
				t.Errorf("error with blue channel: got '%v', want '%v'", gotB, tc.want.Blue())
			}

			if gotA != tc.want.Alpha() {
				t.Errorf("error with alpha factor: got '%v', want '%v'", gotA, tc.want.Alpha())
			}
		})
	}

}

func TestFrameReader(t *testing.T) {
	testCases := []struct {
		name        string
		compression CompressionType
		computation func(r io.Reader) (uint32, uint32, uint32, uint32, error)
		filepath    string
		want        Frame
		wantErr     bool
	}{
		{
			name:        "successful-average-creation",
			compression: MPEG2,
			computation: Average,
			filepath:    "test1.png",
			want: Frame{
				r: 7903,
				g: 7824,
				b: 6843,
				a: 9610,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fr := FrameReader{
				Compression: tc.compression,
				Compute:     tc.computation,
			}
			reader, err := os.Open(tc.filepath)
			if err != nil {
				t.Fatalf("error opening file: err = %v", err)
			}
			defer reader.Close()
			gotR, gotG, gotB, gotA, err := fr.Process(reader)
			if err != nil {
				if tc.wantErr {
					return
				}
				t.Fatalf("unexpected error: got '%v', want 'nil'", err)
			}
			if gotR != tc.want.Red() {
				t.Errorf("error with red channel: got '%v', want '%v'", gotR, tc.want.Red())
			}

			if gotG != tc.want.Green() {
				t.Errorf("error with green channel: got '%v', want '%v'", gotG, tc.want.Green())
			}

			if gotB != tc.want.Blue() {
				t.Errorf("error with blue channel: got '%v', want '%v'", gotB, tc.want.Blue())
			}

			if gotA != tc.want.Alpha() {
				t.Errorf("error with alpha factor: got '%v', want '%v'", gotA, tc.want.Alpha())
			}
		})
	}
}

func TestFrameService(t *testing.T) {
	testCases := []struct {
		name     string
		reader   FrameReader
		filepath string
		want     Frame
		wantErr  bool
	}{
		{
			name: "successful-average-creation",
			reader: FrameReader{
				Compression: MPEG2,
				Compute:     Average,
			},
			filepath: "test1.png",
			want: Frame{
				r: 7903,
				g: 7824,
				b: 6843,
				a: 9610,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fs := NewFrameService(tc.reader)
			reader, err := os.Open(tc.filepath)
			if err != nil {
				t.Fatalf("error opening file: err = %v", err)
			}
			defer reader.Close()
			gotR, gotG, gotB, gotA, err := fs.Read(reader)
			if err != nil {
				if tc.wantErr {
					return
				}
				t.Fatalf("unexpected error: got '%v', want 'nil'", err)
			}
			if gotR != tc.want.Red() {
				t.Errorf("error with red channel: got '%v', want '%v'", gotR, tc.want.Red())
			}

			if gotG != tc.want.Green() {
				t.Errorf("error with green channel: got '%v', want '%v'", gotG, tc.want.Green())
			}

			if gotB != tc.want.Blue() {
				t.Errorf("error with blue channel: got '%v', want '%v'", gotB, tc.want.Blue())
			}

			if gotA != tc.want.Alpha() {
				t.Errorf("error with alpha factor: got '%v', want '%v'", gotA, tc.want.Alpha())
			}
		})
	}
}
