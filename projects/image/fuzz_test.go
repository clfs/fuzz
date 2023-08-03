// Package image_fuzz contains fuzz tests for golang.org/x/image.
package image_fuzz

import (
	"bytes"
	"image"
	"testing"

	"golang.org/x/image/bmp"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/webp"
)

// FuzzBMPDecode ensures that BMP decoding is robust.
func FuzzBMPDecode(f *testing.F) {
	var buf bytes.Buffer
	bmp.Encode(&buf, image.NewRGBA(image.Rect(0, 0, 1, 1)))
	f.Add(buf.Bytes())

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
