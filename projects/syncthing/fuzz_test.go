package syncthing_fuzz

import (
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/syncthing/syncthing/lib/fs"
)

// FuzzSanitizePath ensures SanitizePath returns printable UTF-8 characters.
func FuzzSanitizePath(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		path := fs.SanitizePath(s)
		if !utf8.ValidString(path) {
			t.Errorf("SanitizePath(%q) => %q, not valid UTF-8", s, path)
			return
		}
		for _, c := range path {
			if !unicode.IsPrint(c) {
				t.Errorf("non-printable rune %q in sanitized path", c)
			}
		}
	})
}
