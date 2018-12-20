package loft

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTask(ctx context.Context, input NewTask) (Task, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateEvent(ctx context.Context, input NewEvent) (Event, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateRequest(ctx context.Context, input NewRequest) (Request, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Loft(ctx context.Context) ([]Loft, error) {
	panic("not implemented")
}
func (r *queryResolver) Echo(ctx context.Context) (*Echo, error) {
	panic("not implemented")
}
