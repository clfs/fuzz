// Package riff contains RIFF fuzzing tests.
package riff

import (
	"bytes"
	"testing"

	"golang.org/x/image/riff"
)

func FuzzReader(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		_, r, err := riff.NewReader(bytes.NewReader(b))
		if err != nil {
			t.Skip()
		}

		for {
			_, _, _, err := r.Next()
			if err != nil {
				t.Skip()
			}
		}
	})
}
