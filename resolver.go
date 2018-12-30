package loft

import (
	"context"
	"time"

	"github.com/louistsaitszho/loft/models"
)

type Resolver struct{}

func (r *Resolver) Loft() LoftResolver {
	return &loftResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type loftResolver struct{ *Resolver }

func (r *loftResolver) Members(ctx context.Context, obj *models.Loft) ([]Member, error) {
	//TODO: sql query: get members by loft id
	panic("not implemented")
}
func (r *loftResolver) Tasks(ctx context.Context, obj *models.Loft) ([]Task, error) {
	//TODO: sql query: get tasks by loft id
	panic("not implemented")
}
func (r *loftResolver) Events(ctx context.Context, obj *models.Loft) ([]Event, error) {
	//TODO: sql query: get events by loft id
	panic("not implemented")
}
func (r *loftResolver) Requests(ctx context.Context, obj *models.Loft) ([]Request, error) {
	//TODO: sql query: get requests by loft id
	panic("not implemented")
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
func (r *mutationResolver) CreateLoft(ctx context.Context, input NewLoft) (models.Loft, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLoftAndMember(ctx context.Context, input *NewLoftNewMember) (LoftAndFirstMember, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Lofts(ctx context.Context) ([]models.Loft, error) {
	panic("not implemented")
}
func (r *queryResolver) Loft(ctx context.Context, id string) (*models.Loft, error) {
	panic("not implemented")
}
func (r *queryResolver) Echo(ctx context.Context) (*Echo, error) {
	echo := Echo{
		Time:   time.Now().UTC().Format(time.RFC3339),
		Format: "RFC3339",
	}
	return &echo, nil
}
