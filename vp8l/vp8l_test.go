// Package vp8l contains VP8L fuzzing tests.
package vp8l

import (
	"bytes"
	"testing"

	"golang.org/x/image/vp8l"
)

func FuzzDecode(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		cfg, err := vp8l.DecodeConfig(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}
		if cfg.Width*cfg.Height > 1e6 {
			t.Skip()
		}
		_, _ = vp8l.Decode(bytes.NewReader(b))
	})
}
