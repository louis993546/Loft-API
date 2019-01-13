package models

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

// Loft removes all the big objects, so that gqlgen will look for additional resolver
// - Members & members count
// - Tasks & tasks count
// - Events & events count
// - Notes & notes count
// - Join requests & join requests count
type Loft struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	JoinCode  string    `json:"join_code"`
	CreatedAt time.Time `json:"created_at"`
}

// Member removes a big member struct, so that gqlgen can only run the query when necessary
type Member struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	ApprovedAt *time.Time `json:"approved_at"`
}

type TaskState string

const (
	TaskStateNotDone TaskState = "NOT_DONE"
	TaskStateDone    TaskState = "DONE"
)

func (e TaskState) IsValid() bool {
	switch e {
	case TaskStateNotDone, TaskStateDone:
		return true
	}
	return false
}

func (e TaskState) String() string {
	return string(e)
}

func (e *TaskState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskState", str)
	}
	return nil
}

func (e TaskState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Task removes the 2 big Member struct, so that gqlgen can use resolver to fetch it only when necessary
type Task struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	State     TaskState  `json:"state"`
	CreatedAt time.Time  `json:"createdAt"`
	DueAt     *time.Time `json:"dueAt"`
}
