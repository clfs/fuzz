// Package hcl contains HCL fuzzing tests.
package hcl

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/json"
)

func FuzzJSONParseWithStartPos(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, l, c int) {
		_, _ = json.ParseWithStartPos(b, "fuzz.hcl", hcl.Pos{Line: l, Column: c})
	})
}

func FuzzParseTemplate(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, l, c int) {
		_, _ = hclsyntax.ParseTemplate(b, "fuzz.hcl", hcl.Pos{Line: l, Column: c})
	})
}

func FuzzParseTraversalAbs(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, l, c int) {
		_, _ = hclsyntax.ParseTraversalAbs(b, "fuzz.hcl", hcl.Pos{Line: l, Column: c})
	})
}

func FuzzParseExpression(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, l, c int) {
		_, _ = hclsyntax.ParseExpression(b, "fuzz.hcl", hcl.Pos{Line: l, Column: c})
	})
}

func FuzzParseConfig(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, l, c int) {
		_, _ = hclsyntax.ParseConfig(b, "fuzz.hcl", hcl.Pos{Line: l, Column: c})
	})
}
