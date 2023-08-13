// Package image_fuzz contains fuzz tests for golang.org/x/image.
package image_fuzz

import (
	"bytes"
	"testing"

	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/webp"
)

// FuzzWEBPDecode ensures that WEBP decoding does not panic.
//
// TODO: Add a seed corpus entry.
func FuzzWEBPDecode(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		cfg, err := webp.DecodeConfig(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}

		if cfg.Width*cfg.Height > 1e6 {
			t.Skip()
		}

		_, _ = webp.Decode(bytes.NewReader(b))
	})
}

// FuzzOpenTypeParse ensures that OpenType font parsing does not panic.
func FuzzOpenTypeParse(f *testing.F) {
	f.Add(gomono.TTF)
	f.Fuzz(func(t *testing.T, b []byte) {
		_, _ = opentype.Parse(b)
	})
}

// TODO: Fuzz VP8 and VP8L decoding.
