// Package vp8 contains VP8 fuzzing tests.
package vp8

import (
	"bytes"
	"testing"

	"golang.org/x/image/vp8"
)

func FuzzDecoder(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		dec := vp8.NewDecoder()
		dec.Init(bytes.NewReader(b), len(b))
		if _, err := dec.DecodeFrameHeader(); err != nil {
			t.Skip()
		}
		_, _ = dec.DecodeFrame()
	})
}
