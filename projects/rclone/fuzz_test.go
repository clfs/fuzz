// Package rclone_fuzz contains fuzz tests for github.com/rclone/rclone.
package rclone_fuzz

import (
	"testing"

	"github.com/rclone/rclone/fs/fspath"
	"github.com/rclone/rclone/lib/encoder/filename"
)

// FuzzFilenameRoundTrip ensures that filename encoding and decoding is robust.
func FuzzFilenameRoundTrip(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		_, _ = filename.Decode(s) // no panic
		enc := filename.Encode(s)
		dec, err := filename.Decode(enc)
		if err != nil {
			t.Fatalf("failed to decode roundtripped filename: %v", err)
		}
		if dec != s {
			t.Fatalf("roundtripped filename changed: %q -> %q -> %q", s, enc, dec)
		}
	})
}

// FuzzFspathParse ensures that fspath.Parse is robust.
func FuzzFspathParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		parsed, err := fspath.Parse(path)
		if err != nil {
			t.Skip()
		}

		if parsed.Name == "" {
			if parsed.ConfigString != "" {
				t.Fatalf("bad ConfigString")
			}
			if parsed.Path != path {
				t.Fatalf("local path not preserved")
			}
		} else {
			if parsed.ConfigString+":"+parsed.Path != path {
				t.Fatalf("didn't split properly")
			}
		}
	})
}
