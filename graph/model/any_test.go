package model_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rokiyama/gqlgen-todos/graph/model"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUnmarshalAny(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    *anypb.Any
		wantErr bool
	}{
		{
			name:  "bool",
			value: true,
			want: &anypb.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.BoolValue",
				Value:   marshal(t, wrapperspb.Bool(true)),
			},
			wantErr: false,
		}, {
			name:  "string",
			value: "test",
			want: &anypb.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.StringValue",
				Value:   marshal(t, wrapperspb.String("test")),
			},
			wantErr: false,
		}, {
			name:    "bytes",
			value:   []byte("test"),
			wantErr: true,
		}, {
			name:    "int",
			value:   math.MaxInt32,
			wantErr: true,
		}, {
			name:    "int32",
			value:   int32(math.MaxInt32),
			wantErr: true,
		}, {
			name:  "int64",
			value: int64(math.MaxInt64),
			want: &anypb.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.Int64Value",
				Value:   marshal(t, wrapperspb.Int64(math.MaxInt64)),
			},
			wantErr: false,
		}, {
			name:  "float",
			value: math.MaxFloat32,
			want: &anypb.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.DoubleValue",
				Value:   marshal(t, wrapperspb.Double(math.MaxFloat32)),
			},
			wantErr: false,
		}, {
			name:    "float32",
			value:   float32(math.MaxFloat32),
			wantErr: true,
		}, {
			name:  "float64",
			value: float64(math.MaxFloat64),
			want: &anypb.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.DoubleValue",
				Value:   marshal(t, wrapperspb.Double(math.MaxFloat64)),
			},
			wantErr: false,
		}, {
			name:  "map",
			value: map[string]interface{}{"key": "value"},
			want: &anypb.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.Struct",
				Value:   marshal(t, structValue(t, map[string]interface{}{"key": "value"})),
			},
			wantErr: false,
		}, {
			name:  "array",
			value: []interface{}{"a", "b", 3},
			want: &anypb.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.ListValue",
				Value:   marshal(t, listValue(t, []interface{}{"a", "b", 3})),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.UnmarshalAny(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalAny() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("UnmarshalAny() output mismatch (-want +got):\n%s", diff)
			}

		})
	}
}

func marshal(t *testing.T, m proto.Message) []byte {
	b, err := proto.MarshalOptions{AllowPartial: true, Deterministic: true}.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func structValue(t *testing.T, m map[string]interface{}) *structpb.Struct {
	t.Helper()
	ret, err := structpb.NewStruct(m)
	if err != nil {
		t.Fatal(err)
	}
	return ret
}

func listValue(t *testing.T, l []interface{}) *structpb.ListValue {
	t.Helper()
	ret, err := structpb.NewList(l)
	if err != nil {
		t.Fatal(err)
	}
	return ret
}
