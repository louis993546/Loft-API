package models

import "time"
import "github.com/gofrs/uuid"

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

// Member removes a big member object, so that gqlgen can only run the query when necessary
type Member struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	ApprovedAt *time.Time `json:"approved_at"`
}
