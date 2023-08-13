package toml

import (
	"bytes"
	"testing"

	"github.com/BurntSushi/toml"
)

func FuzzDecode_BurntSushi(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var v any
		_, err := toml.NewDecoder(bytes.NewReader(b)).Decode(&v)
		if err != nil {
			t.Skip()
		}

		var buf bytes.Buffer
		err = toml.NewEncoder(&buf).Encode(v)
		if err != nil {
			t.Errorf("failed to encode valid toml: %v", err)
		}

		var v2 any
		_, err = toml.Decode(buf.String(), &v2)
		if err != nil {
			t.Errorf("failed to decode round-tripped toml: %v", err)
		}
	})
}
