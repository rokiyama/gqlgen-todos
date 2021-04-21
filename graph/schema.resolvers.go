package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/pkg/errors"
	"github.com/rokiyama/gqlgen-todos/graph/generated"
	"github.com/rokiyama/gqlgen-todos/graph/model"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context, criteria *anypb.Any) ([]*model.Todo, error) {
	if criteria != nil {
		v, err := criteria.UnmarshalNew()
		if err != nil {
			log.Printf("unmarshal err: %v", err)
			return nil, err
		}
		log.Printf("criteria: typeUrl=%v; value=%v; type of value=%T", criteria.TypeUrl, v, v)
		switch v := v.(type) {
		case *wrapperspb.Int32Value:
			log.Printf("int32: %d", v.GetValue())
		case *wrapperspb.Int64Value:
			log.Printf("int64: %d", v.GetValue())
		case *wrapperspb.FloatValue:
			log.Printf("float: %f", v.GetValue())
		case *wrapperspb.DoubleValue:
			log.Printf("double: %f", v.GetValue())
		case *wrapperspb.StringValue:
			log.Printf("string: %s", v.GetValue())
		case *wrapperspb.BoolValue:
			log.Printf("bool: %t", v.GetValue())
		case *structpb.Struct:
			log.Printf("map: %v", v.AsMap())
		case *structpb.ListValue:
			log.Printf("list: %v", v.AsSlice())
		default:
			err := errors.Errorf("error: typeUrl=%v; value=%v; type of value=%T", criteria.TypeUrl, v, v)
			log.Print(err)
			return nil, err
		}
	}
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *todoResolver) Any(ctx context.Context, obj *model.Todo) (*anypb.Any, error) {
	panic(fmt.Errorf("not implemented"))
}
