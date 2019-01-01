package models

// Loft is basically what the database can retrieve in a reasonable manner
// TODO: I think even the count don't belong here. they probably should have their own resolvers
type Loft struct {
	ID            string //TODO: use uuid instead?
	Name          string
	MembersCount  int
	TasksCount    int
	EventsCount   int
	RequestsCount int
}
