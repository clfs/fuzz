// Package protobufold_fuzz contains fuzz tests for github.com/golang/protobuf.
package protobufold_fuzz

import (
	"testing"

	"github.com/golang/protobuf/proto" //lint:ignore SA1019 fuzzing despite deprecation
	anypb "github.com/golang/protobuf/ptypes/any"
)

// FuzzUnmarshal ensures proto.Unmarshal does not panic.
func FuzzUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var pb anypb.Any
		_ = proto.Unmarshal(b, &pb)
	})
}
