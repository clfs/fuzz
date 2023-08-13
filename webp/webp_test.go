// Package webp contains WEBP fuzzing tests.
package webp

import (
	"bytes"
	"testing"

	"golang.org/x/image/webp"
)

func FuzzDecode(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		cfg, err := webp.DecodeConfig(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}

		if cfg.Width*cfg.Height > 1e6 {
			t.Skip()
		}

		image, err := webp.Decode(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}

		if image.Bounds().Dx() != cfg.Width || image.Bounds().Dy() != cfg.Height {
			t.Error("bounds do not match config")
		}
	})
}
