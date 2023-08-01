package fuzz

import (
	"bytes"
	"testing"

	btoml "github.com/BurntSushi/toml"
	ptoml "github.com/pelletier/go-toml/v2"
)

func FuzzTOML_BurntSushi(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var v any
		_, err := btoml.NewDecoder(bytes.NewReader(data)).Decode(&v)
		if err != nil {
			t.Skip()
		}

		var buf bytes.Buffer
		if err := btoml.NewEncoder(&buf).Encode(v); err != nil {
			t.Errorf("failed to re-encode decoded toml: %v", err)
		}
	})
}

func FuzzTOML_pelletier(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte, a, b, c, d bool, s string) {
		dec := ptoml.NewDecoder(bytes.NewReader(data))
		if a {
			dec.DisallowUnknownFields()
		}

		var v any
		if err := dec.Decode(&v); err != nil {
			t.Skip()
		}

		var buf bytes.Buffer
		enc := ptoml.NewEncoder(&buf).SetArraysMultiline(b).SetIndentSymbol(s).SetIndentTables(c).SetTablesInline(d)
		if err := enc.Encode(v); err != nil {
			t.Errorf("failed to re-encode decoded toml: %v", err)
		}
	})
}
