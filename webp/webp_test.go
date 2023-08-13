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

		w, h := image.Bounds().Dx(), image.Bounds().Dy()
		if cfg.Width != w || cfg.Height != h {
			t.Errorf("decoded image size mismatch: want %dx%d, got %dx%d", cfg.Width, cfg.Height, w, h)
		}
	})
}
