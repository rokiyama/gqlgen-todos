package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rokiyama/gqlgen-todos/graph/generated"
	"github.com/rokiyama/gqlgen-todos/graph/model"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (r *queryResolver) Test(ctx context.Context, criteria *anypb.Any) (*model.Result, error) {
	if criteria != nil {
		v, err := criteria.UnmarshalNew()
		if err != nil {
			return nil, err
		}
		var vs string
		switch v := v.(type) {
		case *wrapperspb.BoolValue:
			vs = fmt.Sprintf("%t", v.GetValue())
		case *wrapperspb.Int64Value:
			vs = fmt.Sprintf("%d", v.GetValue())
		case *wrapperspb.DoubleValue:
			vs = fmt.Sprintf("%f", v.GetValue())
		case *wrapperspb.StringValue:
			vs = v.GetValue()
		case *structpb.Struct:
			b, err := json.Marshal(v.AsMap())
			if err != nil {
				return nil, err
			}
			vs = string(b)
		case *structpb.ListValue:
			b, err := json.Marshal(v.AsSlice())
			if err != nil {
				return nil, err
			}
			vs = string(b)
		default:
			return nil, fmt.Errorf("invalid type: %T", v)
		}
		return &model.Result{
			TypeURL:   criteria.TypeUrl,
			Value:     vs,
			ValueType: fmt.Sprintf("%T", v),
		}, nil
	}
	return &model.Result{
		TypeURL:   "",
		Value:     fmt.Sprintf("%v", criteria),
		ValueType: fmt.Sprintf("%T", criteria),
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
