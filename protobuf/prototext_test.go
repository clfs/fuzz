package protobuf

import (
	"testing"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/known/anypb"
)

func FuzzPrototextUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var pb anypb.Any
		if err := prototext.Unmarshal(b, &pb); err != nil {
			t.Skip()
		}

		_ = prototext.Format(&pb) // no panic

		b2, err := prototext.Marshal(&pb)
		if err != nil {
			t.Errorf("failed to marshal valid pb: %v", err)
		}
		if err := prototext.Unmarshal(b2, &pb); err != nil {
			t.Errorf("failed to unmarshal round-tripped pb: %v", err)
		}

		_ = prototext.Format(&pb) // no panic
	})
}
