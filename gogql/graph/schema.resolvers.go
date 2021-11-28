package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/golangproj/oofp/gogql/graph/generated"
	"github.com/golangproj/oofp/gogql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return AddNewTodo(input), nil
}

func (r *mutationResolver) RemoveTodo(ctx context.Context, id string) (*model.Todo, error) {
	return RemoveTodo(id)
}

func (r *mutationResolver) SetTodoDone(ctx context.Context, id string, done *bool) (*model.Todo, error) {
	return SetTodoDone(id, done)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return GetAllTodos(), nil
}

func (r *subscriptionResolver) TodoChanges(ctx context.Context) (<-chan *model.TodoEvent, error) {
	return get_subchannel(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
