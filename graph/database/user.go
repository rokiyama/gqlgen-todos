package database

import "github.com/rokiyama/gqlgen-todos/graph/model"

func GetUsers(ids ...string) ([]*model.User, error) {
	var ret []*model.User
	for _, id := range ids {
		ret = append(ret, &model.User{ID: id, Name: "user " + id})
	}
	return ret, nil
}
