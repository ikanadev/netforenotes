package app

import (
	"github.com/oklog/ulid/v2"
	"github.com/vmkevv/netforenotes/interface/api/openapi"
)

type Repository interface {
	SaveNote(note Note)
	UpdateNote(note Note)
	GetNote(id ulid.ULID) (*Note, error)
	DeleteNote(id ulid.ULID)
	GetNotes(params openapi.GetNotesParams) []Note
}
