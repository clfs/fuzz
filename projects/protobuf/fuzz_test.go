// Package protobuf_fuzz contains fuzz tests for google.golang.org/protobuf.
package protobuf_fuzz

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// FuzzUnmarshal ensures proto.Unmarshal does not panic.
func FuzzUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var pb anypb.Any
		_ = proto.Unmarshal(b, &pb)
	})
}
