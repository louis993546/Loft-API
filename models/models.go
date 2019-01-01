package models

import "time"
import "github.com/gofrs/uuid"

// Loft is basically what the database can retrieve in a reasonable manner
type Loft struct {
	ID        uuid.UUID
	Name      string
	JoinCode  string
	CreatedAt time.Time
}
