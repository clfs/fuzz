// Package yaml contains YAML fuzzing tests.
package yaml

import (
	"bytes"
	"testing"

	"gopkg.in/yaml.v3"
)

// TODO: Add github.com/goccy/go-yaml fuzz tests that don't immediately panic.

func FuzzDecode(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var v any
		err := yaml.NewDecoder(bytes.NewReader(b)).Decode(&v)
		if err != nil {
			t.Skip()
		}

		var buf bytes.Buffer
		err = yaml.NewEncoder(&buf).Encode(&v)
		if err != nil {
			t.Errorf("failed to encode valid YAML: %v", err)
		}

		var v2 any
		err = yaml.NewDecoder(&buf).Decode(&v2)
		if err != nil {
			t.Errorf("failed to decode round-tripped YAML: %v", err)
		}
	})
}
