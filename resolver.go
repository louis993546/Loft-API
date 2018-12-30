package loft

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/louistsaitszho/loft/models"
)

// Resolver is the entry point to how the query tree got processed
type Resolver struct {
	db *sql.DB
}

// NewResolver is essentially the constructor for Resolver. It reminds user that they should give Resolver a db to access
func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{
		db: db,
	}
}

// Loft returns a resolver that is able to resolve struct Loft
func (r *Resolver) Loft() LoftResolver {
	return &loftResolver{r}
}

// Mutation returns a resolver that is able to resolve mutations
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns a resolver that is able to resolve query
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type loftResolver struct{ *Resolver }

func (r *loftResolver) Members(ctx context.Context, obj *models.Loft) ([]Member, error) {
	panic("not implemented")
}
func (r *loftResolver) Tasks(ctx context.Context, obj *models.Loft) ([]Task, error) {
	panic("not implemented")
}
func (r *loftResolver) Events(ctx context.Context, obj *models.Loft) ([]Event, error) {
	panic("not implemented")
}
func (r *loftResolver) Requests(ctx context.Context, obj *models.Loft) ([]Request, error) {
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
func (r *queryResolver) Echo(ctx context.Context) (Echo, error) {
	dbStats := r.db.Stats()
	log.Printf("%+v\n", dbStats)

	echo := Echo{
		Time:   time.Now().UTC().Format(time.RFC3339),
		Format: "RFC3339",
	}
	return echo, nil
}
