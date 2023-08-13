// Package image_fuzz contains fuzz tests for golang.org/x/image.
package image_fuzz

import (
	"testing"

	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/opentype"
)

// FuzzOpenTypeParse ensures that OpenType font parsing does not panic.
func FuzzOpenTypeParse(f *testing.F) {
	f.Add(gomono.TTF)
	f.Fuzz(func(t *testing.T, b []byte) {
		_, _ = opentype.Parse(b)
	})
}

// TODO: Fuzz VP8 and VP8L decoding.
