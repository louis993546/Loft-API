package loft

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/louistsaitszho/loft/models"
)

// Resolver is the entry point to how the query tree got processed. It has all the prepared statements that the api would need.
// dbAreYouSureAboutThis is temporary only. All queries should be prepared on server start-up (fail-fast)
type Resolver struct {
	dbAreYouSureAboutThis *sql.DB
	memberCountStmt       *sql.Stmt
	membersStmt           *sql.Stmt
	taskCountStmt         *sql.Stmt
	tasksStmt             *sql.Stmt
	eventCountStmt        *sql.Stmt
	eventsStmt            *sql.Stmt
	noteCountStml         *sql.Stmt
	notesStmt             *sql.Stmt
	loftStmt              *sql.Stmt
	joinRequestCountStmt  *sql.Stmt
	joinRequestStmt       *sql.Stmt
}

// NewResolver is essentially the constructor for Resolver. It reminds user that they should give Resolver a db to access
func NewResolver(db *sql.DB) *Resolver {
	memberCountStmt, mcErr := db.Prepare("SELECT COUNT(*) FROM loft.member WHERE loft.member.loft_id = $1;")
	if mcErr != nil {
		log.Panicf("Invalid query for member count: '%v'\n", mcErr)
	}
	membersStmt, mErr := db.Prepare("SELECT m.id, m.name, m.approved_at, m.join_request_id FROM loft.member AS m WHERE m.loft_id = $1;")
	if mErr != nil {
		log.Panicf("Invalid query for members: '%v'\n", mErr)
	}
	taskCountStmt, tcErr := db.Prepare("SELECT COUNT(*) FROM loft.task WHERE loft.task.loft_id = $1;")
	if tcErr != nil {
		log.Panicf("Invalid query for task count: '%v'\n", tcErr)
	}
	tasksStmt, tErr := db.Prepare("SELECT t.id, t.creator_id, t.created_at, t.assignee_id, t.title, t.due_date FROM loft.task AS t WHERE t.loft_id = $1;")
	if tErr != nil {
		log.Panicf("Invalid query for tasks: '%v'\n", tErr)
	}
	eventCountStmt, ecErr := db.Prepare("SELECT COUNT(*) FROM loft.event WHERE loft.event.loft_id = $1;")
	if ecErr != nil {
		log.Panicf("Invalid query for event count: '%v'\n", ecErr)
	}
	eventsStmt, eErr := db.Prepare("SELECT e.id, e.creator_id, e.created_at, e.start_time, e.end_time, e.title FROM loft.event AS e WHERE e.loft_id = $1;")
	if eErr != nil {
		log.Panicf("Invalid query for events: '%v'\n", eErr)
	}
	noteCountStmt, ncErr := db.Prepare("SELECT COUNT(*) FROM loft.note WHERE loft.note.loft_id = $1;")
	if ncErr != nil {
		log.Panicf("Invalid query for note count: '%v'\n", ncErr)
	}
	//TODO: this should JOIN with SELECT * FROM loft.note_format to get directly the format name
	notesStmt, nErr := db.Prepare("SELECT n.id, n.creator_id, n.created_at, n.format, n.content FROM loft.note AS n WHERE n.loft_id = $1;")
	if nErr != nil {
		log.Panicf("Invalid query for note: '%v'\n", nErr)
	}
	loftStmt, lErr := db.Prepare("SELECT l.id, l.name, l.join_code, l.created_at FROM loft.loft AS l WHERE l.id = $1;")
	if lErr != nil {
		log.Panicf("Invalid query for loft: '%v'\n", lErr)
	}
	joinRequestCountStmt, jrcErr := db.Prepare("SELECT COUNT(*) FROM loft.join_request WHERE loft.join_request.loft_id = $1;")
	if jrcErr != nil {
		log.Panicf("Invalid query for join request count: '%v'\n", jrcErr)
	}
	joinRequestStmt, jrErr := db.Prepare("SELECT jr.id, jr.name, jr.message, jr.created_at FROM loft.join_request AS jr WHERE jr.loft_id = $1;")
	if jrErr != nil {
		log.Panicf("Invalid query for join requests: '%v'\n", jrErr)
	}

	return &Resolver{
		dbAreYouSureAboutThis: db,
		memberCountStmt:       memberCountStmt,
		membersStmt:           membersStmt,
		taskCountStmt:         taskCountStmt,
		tasksStmt:             tasksStmt,
		eventCountStmt:        eventCountStmt,
		eventsStmt:            eventsStmt,
		noteCountStml:         noteCountStmt,
		notesStmt:             notesStmt,
		loftStmt:              loftStmt,
		joinRequestCountStmt:  joinRequestCountStmt,
		joinRequestStmt:       joinRequestStmt,
	}
}

// Event returns a resolver that is able to resolve models.Event struct
func (r *Resolver) Event() EventResolver {
	return &eventResolver{r}
}

// Loft returns a resolver that is able to resolve models.Loft struct
func (r *Resolver) Loft() LoftResolver {
	return &loftResolver{r}
}

// Member returns a resolver that is able to resolve models.Member struct
func (r *Resolver) Member() MemberResolver {
	return &memberResolver{r}
}

// Mutation returns a resolver that is able to resolve mutations
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns a resolver that is able to resolve query
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Task returns a resolver that is able to resolve models.Task struct
func (r *Resolver) Task() TaskResolver {
	return &taskResolver{r}
}

type eventResolver struct{ *Resolver }

func (r *eventResolver) Creator(ctx context.Context, obj *models.Event) (models.Member, error) {
	panic("not implemented")
}

type loftResolver struct{ *Resolver }

func (r *loftResolver) MembersCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.memberCountStmt.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) Members(ctx context.Context, obj *models.Loft) ([]models.Member, error) {
	rows, queryError := r.membersStmt.Query(obj.ID)
	if queryError != nil {
		panic("not implemented: maybe 404, maybe 5XX, etc")
	}
	defer rows.Close()

	var members []models.Member
	for rows.Next() {
		var (
			id            uuid.UUID
			name          string
			approvedAt    string
			joinRequestID uuid.UUID
		)

		if scanErr := rows.Scan(&id, &name, &approvedAt, &joinRequestID); scanErr != nil {
			panic("not implemented")
		}
		members = append(members, models.Member{
			ID:   id,
			Name: name,
		})
	}
	return members, nil
}
func (r *loftResolver) TasksCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.taskCountStmt.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) Tasks(ctx context.Context, obj *models.Loft) ([]models.Task, error) {
	rows, queryError := r.tasksStmt.Query(obj.ID)
	if queryError != nil {
		panic("not implemented")
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var (
			id        uuid.UUID
			title     string
			taskState models.TaskState
			createdAt time.Time
			dueAt     *time.Time
		)

		if scanErr := rows.Scan(&id, &title, &taskState, &createdAt, dueAt); scanErr != nil {
			panic("not implemented")
		}
		tasks = append(tasks, models.Task{
			ID:        id,
			Title:     title,
			State:     taskState,
			CreatedAt: createdAt,
			DueAt:     dueAt,
		})
	}
	return tasks, nil
}
func (r *loftResolver) EventsCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.eventCountStmt.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) Events(ctx context.Context, obj *models.Loft) ([]models.Event, error) {
	rows, queryError := r.eventsStmt.Query(obj.ID)
	if queryError != nil {
		panic("not implemented")
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var (
			id        uuid.UUID
			title     string
			createdAt time.Time
			startTime *time.Time
			endTime   *time.Time
		)

		if scanErr := rows.Scan(&id, &title, &createdAt, startTime, endTime); scanErr != nil {
			panic("not implemented")
		}
		events = append(events, models.Event{
			ID:        id,
			Title:     title,
			CreatedAt: createdAt,
			StartTime: startTime,
			EndTime:   endTime,
		})
	}
	return events, nil
}
func (r *loftResolver) Notes(ctx context.Context, obj *models.Loft) ([]Note, error) {
	panic("not implemented")
}
func (r *loftResolver) NotesCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.noteCountStml.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) JoinRequestsCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.joinRequestCountStmt.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) JoinRequests(ctx context.Context, obj *models.Loft) ([]JoinRequest, error) {
	panic("not implemented")
}

type memberResolver struct{ *Resolver }

func (r *memberResolver) ApprovedBy(ctx context.Context, obj *models.Member) (*models.Member, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTask(ctx context.Context, input NewTask) (models.Task, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateEvent(ctx context.Context, input NewEvent) (models.Event, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateRequest(ctx context.Context, input NewRequest) (JoinRequest, error) {
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
			id              uuid.UUID
			name            string
			joinCode        string
			createdAtString string //TODO: parse it to time instead?
		)
		if scanErr := rows.Scan(&id, &name, &joinCode, &createdAtString); scanErr != nil {
			panic("not implemented: should i skip this one or go straight to nil, error?")
		}
		lofts = append(lofts, models.Loft{
			ID:       id,
			Name:     name,
			JoinCode: joinCode,
			//TODO: created at  missing
		})
	}

	return lofts, nil
}
func (r *queryResolver) Loft(ctx context.Context, id uuid.UUID) (*models.Loft, error) {
	panic("not implemented")
}
func (r *queryResolver) Echo(ctx context.Context) (Echo, error) {
	echo := Echo{
		Time:   time.Now().UTC().Format(time.RFC3339),
		Format: "RFC3339",
	}
	return echo, nil
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) Creator(ctx context.Context, obj *models.Task) (models.Member, error) {
	panic("not implemented")
}
func (r *taskResolver) Assignee(ctx context.Context, obj *models.Task) (*models.Member, error) {
	panic("not implemented")
}
