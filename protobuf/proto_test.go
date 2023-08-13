package protobuf

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func FuzzProtoUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var pb anypb.Any
		if err := proto.Unmarshal(b, &pb); err != nil {
			t.Skip()
		}
		b2, err := proto.Marshal(&pb)
		if err != nil {
			t.Errorf("failed to marshal valid pb: %v", err)
		}
		if err := proto.Unmarshal(b2, &pb); err != nil {
			t.Errorf("failed to unmarshal round-tripped pb: %v", err)
		}
	})
}
