package database

import (
	"context"
	"log"

	"github.com/graph-gophers/dataloader"
	"github.com/rokiyama/gqlgen-todos/graph/model"
)

func GetUsers(keys dataloader.Keys) ([]*model.User, error) {
	var ret []*model.User
	for _, key := range keys {
		ret = append(ret, &model.User{ID: key.String(), Name: "user " + key.String()})
	}
	return ret, nil
}

func NewLoader() *dataloader.Loader {
	// setup batch function
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		log.Printf("loader: keys=%v", keys)
		users, err := GetUsers(keys)
		if err != nil {
			return []*dataloader.Result{{Error: err}}
		}
		var ret []*dataloader.Result
		for _, u := range users {
			ret = append(ret, &dataloader.Result{
				Data: u,
			})
		}
		return ret
	}
	// create Loader with an in-memory cache
	loader := dataloader.NewBatchedLoader(batchFn)
	return loader
}
