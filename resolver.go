//go:generate go run ./scripts/gqlgen.go -v
package loft

import (
	"context"
	"time"
)

type Resolver struct {
	lofts []Loft // TODO: temporary
}

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

func (r *queryResolver) Lofts(ctx context.Context) ([]Loft, error) {
	return r.lofts, nil
}
func (r *queryResolver) Loft(ctx context.Context, id string) (Loft, error) {
	panic("not implemented")
}
func (r *queryResolver) Echo(ctx context.Context) (*Echo, error) {
	echo := Echo{
		Time:   time.Now().UTC().Format(time.RFC3339),
		Format: "RFC3339",
	}
	return &echo, nil
}
