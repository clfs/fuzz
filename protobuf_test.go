package fuzz

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func FuzzProtobuf(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b, c, d, e bool, f int, data []byte) {
		uOpts := proto.UnmarshalOptions{
			Merge:          a,
			AllowPartial:   b,
			DiscardUnknown: c,
			RecursionLimit: f,
		}

		mOpts := proto.MarshalOptions{
			AllowPartial:  d,
			Deterministic: e,
		}

		var pb anypb.Any
		if err := uOpts.Unmarshal(data, &pb); err != nil {
			t.Skip()
		}

		_, err := mOpts.Marshal(&pb)
		if err != nil {
			t.Errorf("failed to marshal after unmarshal: %v", err)
		}
	})
}

func FuzzProtobufJSON(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b, c, d, e, f, g bool, h string, data []byte) {
		uOpts := protojson.UnmarshalOptions{
			AllowPartial:   a,
			DiscardUnknown: b,
		}

		mOpts := protojson.MarshalOptions{
			Multiline:       c,
			Indent:          h,
			AllowPartial:    d,
			UseProtoNames:   e,
			UseEnumNumbers:  f,
			EmitUnpopulated: g,
		}

		var pb anypb.Any
		if err := uOpts.Unmarshal(data, &pb); err != nil {
			t.Skip()
		}

		_, err := mOpts.Marshal(&pb)
		if err != nil {
			t.Errorf("failed to marshal after unmarshal: %v", err)
		}

		_ = mOpts.Format(&pb) // no panic
	})
}

func FuzzProtobufText(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b, c, d, e, f bool, g string, data []byte) {
		uOpts := prototext.UnmarshalOptions{
			AllowPartial:   a,
			DiscardUnknown: b,
		}

		mOpts := prototext.MarshalOptions{
			Multiline:    c,
			Indent:       g,
			EmitASCII:    d,
			AllowPartial: e,
			EmitUnknown:  f,
		}

		var pb anypb.Any
		if err := uOpts.Unmarshal(data, &pb); err != nil {
			t.Skip()
		}

		_, err := mOpts.Marshal(&pb)
		if err != nil {
			t.Errorf("failed to marshal after unmarshal: %v", err)
		}

		_ = mOpts.Format(&pb) // no panic
	})
}
