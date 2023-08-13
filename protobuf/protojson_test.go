package protobuf

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
)

func FuzzProtojsonUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		var pb anypb.Any
		if err := protojson.Unmarshal(b, &pb); err != nil {
			t.Skip()
		}

		_ = protojson.Format(&pb) // no panic

		b2, err := protojson.Marshal(&pb)
		if err != nil {
			t.Errorf("failed to marshal valid pb: %v", err)
		}
		if err := protojson.Unmarshal(b2, &pb); err != nil {
			t.Errorf("failed to unmarshal round-tripped pb: %v", err)
		}

		_ = protojson.Format(&pb) // no panic
	})
}
