// Package ccitt contains CCITT fuzzing tests.
package ccitt

import (
	"bytes"
	"image"
	"testing"

	"golang.org/x/image/ccitt"
)

func FuzzDecodeIntoGray(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, w, h int) {
		for _, cfg := range []struct {
			order ccitt.Order
			sf    ccitt.SubFormat
			opts  ccitt.Options
		}{
			{ccitt.LSB, ccitt.Group3, ccitt.Options{Align: true, Invert: true}},
			{ccitt.LSB, ccitt.Group3, ccitt.Options{Align: true, Invert: false}},
			{ccitt.LSB, ccitt.Group3, ccitt.Options{Align: false, Invert: true}},
			{ccitt.LSB, ccitt.Group3, ccitt.Options{Align: false, Invert: false}},
			{ccitt.LSB, ccitt.Group4, ccitt.Options{Align: true, Invert: true}},
			{ccitt.LSB, ccitt.Group4, ccitt.Options{Align: true, Invert: false}},
			{ccitt.LSB, ccitt.Group4, ccitt.Options{Align: false, Invert: true}},
			{ccitt.LSB, ccitt.Group4, ccitt.Options{Align: false, Invert: false}},
			{ccitt.MSB, ccitt.Group3, ccitt.Options{Align: true, Invert: true}},
			{ccitt.MSB, ccitt.Group3, ccitt.Options{Align: true, Invert: false}},
			{ccitt.MSB, ccitt.Group3, ccitt.Options{Align: false, Invert: true}},
			{ccitt.MSB, ccitt.Group3, ccitt.Options{Align: false, Invert: false}},
			{ccitt.MSB, ccitt.Group4, ccitt.Options{Align: true, Invert: true}},
			{ccitt.MSB, ccitt.Group4, ccitt.Options{Align: true, Invert: false}},
			{ccitt.MSB, ccitt.Group4, ccitt.Options{Align: false, Invert: true}},
			{ccitt.MSB, ccitt.Group4, ccitt.Options{Align: false, Invert: false}},
		} {
			r := ccitt.NewReader(bytes.NewReader(b), cfg.order, cfg.sf, w, h, &cfg.opts)
			img := image.NewGray(image.Rect(0, 0, w, h))
			if err := ccitt.DecodeIntoGray(img, r, cfg.order, cfg.sf, &cfg.opts); err != nil {
				t.Skip()
			}

			// Verify bounds.
			gotW, gotH := img.Bounds().Dx(), img.Bounds().Dy()
			if gotW != w || gotH != h {
				t.Errorf("got %dx%d; want %dx%d", gotW, gotH, w, h)
			}
		}
	})
}
