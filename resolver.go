package loft

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/louistsaitszho/loft/models"
)

// Resolver is the entry point to how the query tree got processed
// dbAreYouSureAboutThis is pretty much temporary only. They should be prepared on server start-up (fail-fast)
type Resolver struct {
	dbAreYouSureAboutThis *sql.DB
	memberCountStmt       *sql.Stmt
	membersStmt           *sql.Stmt
	taskCountStmt         *sql.Stmt
	tasksStmt             *sql.Stmt
	eventCountStmt        *sql.Stmt
	eventsStmt            *sql.Stmt
}

// NewResolver is essentially the constructor for Resolver. It reminds user that they should give Resolver a db to access
// TODO: unless it's COUNT don't use *
func NewResolver(db *sql.DB) *Resolver {
	memberCountStmt, mcErr := db.Prepare("SELECT COUNT(*) FROM loft.member WHERE loft.member.loft_id=?;")
	if mcErr != nil {
		log.Panicf("Invalid query for member count: %e\n", mcErr)
	}
	membersStmt, mErr := db.Prepare("SELECT * FROM loft.member WHERE loft.member.loft_id=?;")
	if mErr != nil {
		log.Panicf("Invalid query for members: %e\n", mErr)
	}
	taskCountStmt, tcErr := db.Prepare("SELECT COUNT(*) FROM loft.task WHERE loft.task.loft_id=?;")
	if tcErr != nil {
		log.Panicf("Invalid query for task count: %e\n", tcErr)
	}
	tasksStmt, tErr := db.Prepare("SELECT * FROM loft.task WHERE loft.task.loft_id=?;")
	if tErr != nil {
		log.Panicf("Invalid query for tasks: %e\n", tErr)
	}
	eventCountStmt, ecErr := db.Prepare("SELECT COUNT(*) FROM loft.event WHERE loft.event.loft_id=?;")
	if ecErr != nil {
		log.Panicf("Invalid query for event count: %e\n", ecErr)
	}
	eventsStmt, eErr := db.Prepare("SELECT * FROM loft.event WHERE loft.event.loft_id=?;")
	if eErr != nil {
		log.Panicf("Invalid query for events: %e\n", eErr)
	}
	return &Resolver{
		dbAreYouSureAboutThis: db,
		memberCountStmt:       memberCountStmt,
		membersStmt:           membersStmt,
		taskCountStmt:         taskCountStmt,
		tasksStmt:             tasksStmt,
		eventCountStmt:        eventCountStmt,
		eventsStmt:            eventsStmt,
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

func (r *loftResolver) ID(ctx context.Context, obj *models.Loft) (string, error) {
	//TODO: replace it with external marshaler (https://gqlgen.com/reference/scalars/)
	return obj.ID.String(), nil
}
func (r *loftResolver) MembersCount(ctx context.Context, obj *models.Loft) (int, error) {
	panic("not implemented")
}
func (r *loftResolver) Members(ctx context.Context, obj *models.Loft) ([]Member, error) {
	panic("not implemented")
}
func (r *loftResolver) TasksCount(ctx context.Context, obj *models.Loft) (int, error) {
	panic("not implemented")
}
func (r *loftResolver) Tasks(ctx context.Context, obj *models.Loft) ([]Task, error) {
	panic("not implemented")
}
func (r *loftResolver) EventsCount(ctx context.Context, obj *models.Loft) (int, error) {
	panic("not implemented")
}
func (r *loftResolver) Events(ctx context.Context, obj *models.Loft) ([]Event, error) {
	panic("not implemented")
}
func (r *loftResolver) RequestsCount(ctx context.Context, obj *models.Loft) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM loft.join_request WHERE loft.join_request.loft_id=%s;", obj.ID.String())
	log.Println(query)
	panic("not implemented")
}
func (r *loftResolver) Requests(ctx context.Context, obj *models.Loft) ([]Request, error) {
	query := fmt.Sprintf("SELECT * FROM loft.join_request WHERE loft.join_request.loft_id=%s;", obj.ID.String())
	log.Println(query)
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
	query := "SELECT loft.loft.id, loft.loft.\"name\", loft.loft.join_code, loft.loft.created_at FROM loft.loft;"
	rows, queryErr := r.dbAreYouSureAboutThis.Query(query)
	if queryErr != nil {
		panic("not implemented: find out what can happen, and what is recoverable and how")
	}
	defer rows.Close()

	var lofts []models.Loft
	for rows.Next() {
		var (
			id              uuid.UUID //TODO: I think uuid package has build-in sql support
			name            string
			joinCode        string
			createdAtString string //TODO: parse it to time instead?
		)
		if scanErr := rows.Scan(&id, &name, &joinCode, &createdAtString); scanErr != nil {
			panic("not implemented: should i skip this one or go straight to nil, error?")
		}
		lofts = append(lofts, models.Loft{
			ID:   id,
			Name: name,
			//TODO: join code and created at  missing
		})
	}

	return lofts, nil
}
func (r *queryResolver) Loft(ctx context.Context, id string) (*models.Loft, error) {
	panic("not implemented")
}
func (r *queryResolver) Echo(ctx context.Context) (Echo, error) {
	echo := Echo{
		Time:   time.Now().UTC().Format(time.RFC3339),
		Format: "RFC3339",
	}
	return echo, nil
}
