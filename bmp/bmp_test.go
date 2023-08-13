// Package bmp contains BMP fuzzing tests.
package bmp

import (
	"bytes"
	"testing"

	"golang.org/x/image/bmp"
)

func FuzzDecode(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		cfg, err := bmp.DecodeConfig(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}

		if cfg.Width*cfg.Height > 1e6 {
			t.Skip()
		}

		img, err := bmp.Decode(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}

		var w bytes.Buffer
		err = bmp.Encode(&w, img)
		if err != nil {
			t.Fatalf("failed to encode valid image: %v", err)
		}

		img1, err := bmp.Decode(&w)
		if err != nil {
			t.Fatalf("failed to decode roundtripped image: %v", err)
		}

		got := img1.Bounds()
		want := img.Bounds()
		if !got.Eq(want) {
			t.Fatalf("roundtripped image bounds have changed, got: %s, want: %s", got, want)
		}
	})
}
