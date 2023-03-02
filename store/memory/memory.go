package memory

import (
	"errors"
	"sort"
	"strings"

	"github.com/oklog/ulid/v2"
	"github.com/vmkevv/netforenotes/app"
	"github.com/vmkevv/netforenotes/interface/api/openapi"
)

var NotFoundErr = errors.New("Note not found")

type NotesStore struct {
	notes map[string]app.Note
}

func NewNotesStore() *NotesStore {
	return &NotesStore{
		notes: make(map[string]app.Note),
	}
}

// ensures NotesStore implements app.Respository
var _ app.Repository = (*NotesStore)(nil)

func (ns NotesStore) GetLen() int {
	return len(ns.notes)
}

func (ns *NotesStore) SaveNote(note app.Note) {
	ns.notes[note.ID().String()] = note
}

func (ns *NotesStore) UpdateNote(note app.Note) {
	_, ok := ns.notes[note.ID().String()]
	if ok {
		ns.notes[note.ID().String()] = note
	}
}

func (ns NotesStore) GetNote(id ulid.ULID) (*app.Note, error) {
	note, ok := ns.notes[id.String()]
	if !ok {
		return nil, NotFoundErr
	}
	return &note, nil
}

func (ns *NotesStore) DeleteNote(id ulid.ULID) {
	delete(ns.notes, id.String())
}

func (ns NotesStore) GetNotes(params openapi.GetNotesParams) []app.Note {
	notes := ns.filterNotes(params.Search)
	sort.Slice(notes, func(i, j int) bool {
		sorted := true
		// sort in asc order
		if params.Sort == openapi.Title &&
			strings.ToLower(notes[i].Title()) > strings.ToLower(notes[j].Title()) {
			sorted = false
		}
		if params.Sort == openapi.Date &&
			notes[i].CreatedAt().After(notes[j].CreatedAt()) {
			sorted = false
		}
		// reverse if order is desc
		if params.Order == openapi.Desc {
			sorted = !sorted
		}
		return sorted
	})
	return notes
}

func (ns NotesStore) filterNotes(termRef *string) []app.Note {
	term := ""
	if termRef != nil {
		term = strings.Trim(*termRef, " ")
		term = strings.ToLower(term)
	}
	notes := make([]app.Note, 0)

	for _, note := range ns.notes {
		// if term is empty, just add and continue
		if len(term) == 0 {
			notes = append(notes, note)
			continue
		}
		if strings.Contains(strings.ToLower(note.Title()), term) ||
			strings.Contains(strings.ToLower(note.Body()), term) {
			notes = append(notes, note)
		}
	}
	return notes
}
