package app

import (
	"errors"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

type Note struct {
	id        ulid.ULID
	title     string
	body      string
	createdAt time.Time
}

func (note Note) ID() ulid.ULID {
	return note.id
}

func (note Note) Title() string {
	return note.title
}

func (note Note) Body() string {
	return note.body
}

func (note Note) CreatedAt() time.Time {
	return note.createdAt
}

func (note *Note) UpdateNote(title, body string) error {
	err := validateTitleAndBody(title, body)
	if err != nil {
		return err
	}
	note.title = title
	note.body = body
	return nil
}

func NewNote(title, body string, milliSecs int64) (Note, error) {
	var note Note
	err := validateTitleAndBody(title, body)
	if err != nil {
		return note, err
	}
	if milliSecs == 0 {
		return note, errors.New("Time can not be zero")
	}
	note.id = ulid.Make()
	note.title = title
	note.body = body
	note.createdAt = time.UnixMilli(milliSecs)
	return note, nil
}

func validateTitleAndBody(title, body string) error {
	if strings.Trim(title, " ") == "" {
		return errors.New("Title can not be empty")
	}
	if strings.Trim(body, " ") == "" {
		return errors.New("Body can not be empty")
	}
	return nil
}
