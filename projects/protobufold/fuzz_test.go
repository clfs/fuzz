// Package protobufold_fuzz contains fuzz tests for github.com/golang/protobuf.
package protobufold_fuzz

import (
	"testing"

	proto_deprecated "github.com/golang/protobuf/proto" //lint:ignore SA1019 fuzzing despite deprecation
	anypb_deprecated "github.com/golang/protobuf/ptypes/any"
)

// FuzzUnmarshal ensures proto.Unmarshal does not panic.
func FuzzUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var pb anypb_deprecated.Any
		_ = proto_deprecated.Unmarshal(b, &pb)
	})
}
