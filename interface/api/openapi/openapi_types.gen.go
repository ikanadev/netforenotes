// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package openapi

// Defines values for GetNotesParamsSort.
const (
	Date  GetNotesParamsSort = "date"
	Title GetNotesParamsSort = "title"
)

// Defines values for GetNotesParamsOrder.
const (
	Asc  GetNotesParamsOrder = "asc"
	Desc GetNotesParamsOrder = "desc"
)

// ErrorMessage defines model for ErrorMessage.
type ErrorMessage struct {
	Message string `json:"message"`
}

// Note defines model for Note.
type Note struct {
	Body      string `json:"body"`
	CreatedAt int64  `json:"created_at"`
	Id        string `json:"id"`
	Title     string `json:"title"`
}

// NoteBase defines model for NoteBase.
type NoteBase struct {
	CreatedAt int64  `json:"created_at"`
	Title     string `json:"title"`
}

// NoteCreation defines model for NoteCreation.
type NoteCreation struct {
	Body      string `json:"body"`
	CreatedAt int64  `json:"created_at"`
	Title     string `json:"title"`
}

// NoteList defines model for NoteList.
type NoteList = []NoteListItem

// NoteListItem defines model for NoteListItem.
type NoteListItem struct {
	CreatedAt int64  `json:"created_at"`
	Id        string `json:"id"`
	Title     string `json:"title"`
}

// GetNotesParams defines parameters for GetNotes.
type GetNotesParams struct {
	Sort  GetNotesParamsSort  `form:"sort" json:"sort"`
	Order GetNotesParamsOrder `form:"order" json:"order"`

	// Search Filter notes with some search text in title or body
	Search *string `form:"search,omitempty" json:"search,omitempty"`
}

// GetNotesParamsSort defines parameters for GetNotes.
type GetNotesParamsSort string

// GetNotesParamsOrder defines parameters for GetNotes.
type GetNotesParamsOrder string

// AddNoteJSONRequestBody defines body for AddNote for application/json ContentType.
type AddNoteJSONRequestBody = NoteCreation

// UpdateNoteJSONRequestBody defines body for UpdateNote for application/json ContentType.
type UpdateNoteJSONRequestBody = Note