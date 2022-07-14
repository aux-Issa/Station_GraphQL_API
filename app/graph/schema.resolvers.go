package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aux-Issa/Station_GraphQL_API/graph/generated"
	"github.com/aux-Issa/Station_GraphQL_API/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) StationByName(ctx context.Context, stationName *string) ([]*model.Station, error) {
	return r.getStationByName(ctx, stationName)
}

func (r *queryResolver) StationByCd(ctx context.Context, stationCd *int) (*model.Station, error) {
	return r.getStationByCD(ctx, stationCd)
}

func (r *stationResolver) BeforeStation(ctx context.Context, obj *model.Station) (*model.Station, error) {
	return r.getBeforeStation(ctx, obj)
}

func (r *stationResolver) AfterStation(ctx context.Context, obj *model.Station) (*model.Station, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stationResolver) TransferStation(ctx context.Context, obj *model.Station) ([]*model.Station, error) {
	return r.getTransferStaion(ctx, obj)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Station returns generated.StationResolver implementation.
func (r *Resolver) Station() generated.StationResolver { return &stationResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type stationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
