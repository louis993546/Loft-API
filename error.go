package main

// ErrorResponseObject is a representation of a generic error object in this API, following JSONAPI specification
type ErrorResponseObject struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Code   string `json:"code"`
}

// ErrorResponse is what JSONAPI wants to see when application wants to return error
type ErrorResponse struct {
	Errors []ErrorResponseObject `json:"errors"`
}

func ResponseError404LoftNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the loft you are specifying",
			Code:   "404_LOFT_NOT_FOUND",
		},
	}
}

func ResponseError404NoteNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the note you are specifying",
			Code:   "404_NOTE_NOT_FOUND",
		},
	}
}

func ResponseError404TaskNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the task you are specifying",
			Code:   "404_TASK_NOT_FOUND",
		},
	}
}

func ResponseError404EventNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the event you are specifying",
			Code:   "404_EVENT_NOT_FOUND",
		},
	}
}

func ResponseError404MemberNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the member you are specifying",
			Code:   "404_Member_NOT_FOUND",
		},
	}
}
