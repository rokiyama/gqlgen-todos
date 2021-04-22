package model

import (
	"encoding/json"
	"io"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// refs https://gqlgen.com/reference/scalars/#custom-scalars-with-third-party-types

func MarshalAny(a *anypb.Any) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		b, err := json.Marshal(a)
		if err != nil {
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Printf("io error: %s", err)
			}
			return
		}
		if _, err := w.Write(b); err != nil {
			log.Printf("io error: %s", err)
		}
	})
}

func UnmarshalAny(v interface{}) (*anypb.Any, error) {
	var m proto.Message
	log.Printf("%T", v)
	switch v := v.(type) {
	case nil:
		return nil, nil
	case bool:
		m = wrapperspb.Bool(v)
	// case int:
	// 	m = wrapperspb.Int64(int64(v))
	// case int32:
	// 	m = wrapperspb.Int32(v)
	case int64:
		m = wrapperspb.Int64(v)
	// case uint:
	// 	m = wrapperspb.Int64(int64(v))
	// case uint32:
	// 	m = wrapperspb.Int64(int64(v))
	// case uint64:
	// 	m = wrapperspb.Int64(int64(v))
	// case float32:
	// 	m = wrapperspb.Float(v)
	case float64:
		m = wrapperspb.Double(v)
	case string:
		m = wrapperspb.String(v)
	// case []byte:
	// 	m = wrapperspb.Bytes(v)
	case map[string]interface{}:
		s, err := structpb.NewStruct(v)
		if err != nil {
			return nil, err
		}
		m = s
	case []interface{}:
		l, err := structpb.NewList(v)
		if err != nil {
			return nil, err
		}
		m = l
	default:
		return nil, errors.Errorf("invalid type: %T", v)
	}
	return anypb.New(m)
}
