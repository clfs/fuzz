package toml

import (
	"bytes"
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func FuzzDecode_pelletier(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var v any
		err := toml.NewDecoder(bytes.NewReader(b)).Decode(&v)
		if err != nil {
			t.Skip()
		}

		var buf bytes.Buffer
		err = toml.NewEncoder(&buf).Encode(v)
		if err != nil {
			t.Errorf("failed to encode valid toml: %v", err)
		}

		var v2 any
		err = toml.NewDecoder(&buf).Decode(&v2)
		if err != nil {
			t.Errorf("failed to decode round-tripped toml: %v", err)
		}
	})
}
