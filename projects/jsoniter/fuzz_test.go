// Package jsoniter_fuzz contains fuzz tests for github.com/json-iterator/go.
package jsoniter_fuzz

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func FuzzUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		for _, typ := range []func() any{
			func() any { return new(any) },
			func() any { return new(map[string]any) },
			func() any { return new([]any) },
		} {
			i := typ()
			if err := jsoniter.Unmarshal(b, i); err != nil {
				return
			}

			encoded, err := jsoniter.Marshal(i)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}

			if err := jsoniter.Unmarshal(encoded, i); err != nil {
				t.Fatalf("failed to roundtrip: %v", err)
			}
		}
	})
}

func FuzzParseBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		for _, cfg := range []jsoniter.API{
			jsoniter.ConfigDefault,
			jsoniter.ConfigFastest,
			jsoniter.ConfigCompatibleWithStandardLibrary,
		} {
			iter := jsoniter.ParseBytes(cfg, b)
			for iter.ReadArray() {
				iter.Skip()
			}
		}
	})
}
