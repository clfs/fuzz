package fuzz

import (
	"bytes"
	"image"
	"testing"

	"golang.org/x/image/tiff"
)

func FuzzTIFF(f *testing.F) {
	var buf bytes.Buffer
	tiff.Encode(&buf, image.NewRGBA(image.Rect(0, 0, 1, 1)), nil)
	f.Add(buf.Bytes())

	f.Fuzz(func(t *testing.T, b []byte) {
		cfg, err := tiff.DecodeConfig(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}

		if cfg.Width*cfg.Height > 1e5 {
			t.Skip()
		}

		img, err := tiff.Decode(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}

		for _, c := range []tiff.CompressionType{
			tiff.Uncompressed,
			tiff.Deflate,
		} {
			var w bytes.Buffer

			err = tiff.Encode(&w, img, &tiff.Options{Compression: c})
			if err != nil {
				t.Fatalf("failed to encode valid image: %v", err)
			}

			img1, err := tiff.Decode(&w)
			if err != nil {
				t.Fatalf("failed to decode roundtripped image: %v", err)
			}

			got := img1.Bounds()
			want := img.Bounds()
			if !got.Eq(want) {
				t.Fatalf("roundtripped image bounds have changed, got: %s, want: %s", got, want)
			}
		}
	})
}
