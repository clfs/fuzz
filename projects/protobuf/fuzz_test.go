package protobuf_fuzz

import (
	"testing"

	proto_deprecated "github.com/golang/protobuf/proto" //lint:ignore SA1019 fuzzing despite deprecation
	anypb_deprecated "github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func FuzzUnmarshalDeprecated(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var pb anypb_deprecated.Any
		_ = proto_deprecated.Unmarshal(b, &pb)
	})
}

func FuzzUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var pb anypb.Any
		_ = proto.Unmarshal(b, &pb)
	})
}
