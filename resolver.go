//go:generate go run ./scripts/gqlgen.go -v
package loft

import (
	"context"
	"fmt"
	"sort"
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
func (r *queryResolver) Loft(ctx context.Context, id string) (*Loft, error) {
	loftPos := sort.Search(len(r.lofts), func(i int) bool {
		return r.lofts[i].ID == id
	})

	if loftPos == len(r.lofts) {
		return nil, fmt.Errorf("loft not found for loft #%s", id)
	}

	return &r.lofts[loftPos], nil
}
func (r *queryResolver) Echo(ctx context.Context) (*Echo, error) {
	echo := Echo{
		Time:   time.Now().UTC().Format(time.RFC3339),
		Format: "RFC3339",
	}
	return &echo, nil
}
