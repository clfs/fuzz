// Package sfnt contains SFNT fuzzing tests.
package sfnt

import (
	"bytes"
	"testing"

	"golang.org/x/image/font/sfnt"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		_, _ = sfnt.Parse(b)
		_, _ = sfnt.ParseReaderAt(bytes.NewReader(b))
		_, _ = sfnt.ParseCollection(b)
		_, _ = sfnt.ParseCollectionReaderAt(bytes.NewReader(b))
	})
}
