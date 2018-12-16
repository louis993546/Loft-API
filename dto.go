package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// TODO: this should not be package main
// This package contains all the DTO for the project
// This whole thing follows JSON API, and it is not as intuitive as normal OO DTOs
// Use the public factory instead of create a new struct directly

// <note>

type attributeOfNote struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Content     string    `json:"content"`
	//TODO: creator: id or member object or not an attribute?
	//TODO: format: enum or struct
}

type dataOfNote struct {
	Type       string          `json:"type"`
	ID         string          `json:"id"`
	Attributes attributeOfNote `json:"attributes"`
}

// </note>
// <task>

type attributeOfTask struct {
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	DueDate   time.Time `json:"due_date"`
	//TODO: creator: id or member object or not an attribute?
	//TODO: assignee_id: id or member object or not an attribute?
}

type dataOfTask struct {
	Type       string          `json:"type"`
	ID         string          `json:"id"`
	Attributes attributeOfTask `json:"attributes"`
}

// </task>
// <member>

type attributeOfMember struct {
	ApprovedAt time.Time `json:"approved_at"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	//TODO: Approved by: member id or member object or not an attribute?
	//TODO: Join reuqest id: do they ever need to know that?
}

type dataOfMember struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes attributeOfMember `json:"attributes"`
}

// </member>
// <event>

type attributeOfEvent struct {
	CreatedAt time.Time `json:"created_at"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Title     string    `json:"title"`
	//TODO: creator: id or member object or not an attribute?
}

type dataOfEvent struct {
	Type       string           `json:"type"`
	ID         string           `json:"id"`
	Attributes attributeOfEvent `json:"attributes"`
}

// </event>
// <loft>

type attributeOfLoft struct {
	Name      string    `json:"name"`
	JoinCode  string    `json:"join_code"`
	CreatedAt time.Time `json:"created_at"`
}

type dataOfLoft struct {
	Type       string          `json:"type"` //this should always be "loft"
	ID         string          `json:"id"`
	Attributes attributeOfLoft `json:"attributes"`
}

type relationshipWithLoft struct {
	Loft dataOfLoft `json:"loft"`
}

// <loft>
// <loft_joining_request>

type attributeOfLoftJoiningRequest struct {
}

type dataOfLoftJoiningRequest struct {
	Type       string                        `json:"type"` //this should always be "loft"
	ID         string                        `json:"id"`
	Attributes attributeOfLoftJoiningRequest `json:"attributes"`
}

// </join_request>

// ---------------------DTO Factories----------------------
// http://www.golangpatterns.info/object-oriented/constructors

type noteDto struct {
	Data         dataOfNote
	Relationship relationshipWithLoft
}

// NewNoteDto is a test fx
func NewNoteDto(id string) *noteDto {
	attribute := new(attributeOfNote)
	attribute.Content = "content"
	attribute.CreatedAt = time.Now()
	attribute.Description = "more parameter in the function"
	attribute.Title = "more parameter in the function"

	data := new(dataOfNote)
	data.ID = id
	data.Type = "note"
	data.Attributes = *attribute

	loftAttribute := new(attributeOfLoft)
	loftAttribute.Name = ""

	loft := new(dataOfLoft)
	loft.ID = ""
	loft.Type = "loft"
	loft.Attributes = *loftAttribute

	relationship := new(relationshipWithLoft)
	relationship.Loft = *loft

	note := new(noteDto)
	note.Data = *data
	note.Relationship = *relationship

	return note
}

// ParseCreateNoteRequest is a test fx
func ParseCreateNoteRequest(sth http.Request) (*noteDto, error) {
	var note noteDto
	err := json.NewDecoder(sth.Body).Decode(&note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}
