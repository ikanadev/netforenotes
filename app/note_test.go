package app_test

import (
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/vmkevv/netforenotes/app"
)

func TestNewNote(t *testing.T) {
	title := "Title"
	body := "Note body text"
	milliSecs := time.Now().UnixMilli()

	t.Run("should create note with no errors", func(t *testing.T) {
		assert := assert.New(t)
		_, err := app.NewNote(title, body, milliSecs)
		assert.Nil(err)
	})

	t.Run("getters should get the correct note info", func(t *testing.T) {
		assert := assert.New(t)
		note, _ := app.NewNote(title, body, milliSecs)
		assert.Equal(title, note.Title())
		assert.Equal(body, note.Body())
		assert.Equal(milliSecs, note.CreatedAt().UnixMilli())
		// Should return a valid ulid ID
		noteId := note.ID()
		_, err := ulid.Parse(noteId.String())
		assert.Nil(err)
	})

	t.Run("should return correct errors", func(t *testing.T) {
		assert := assert.New(t)
		_, err := app.NewNote("  ", body, milliSecs)
		assert.EqualError(err, app.EmptyTitleErr.Error())
		_, err = app.NewNote(title, "  ", milliSecs)
		assert.EqualError(err, app.EmptyBodyErr.Error())
		_, err = app.NewNote(title, body, 0)
		assert.EqualError(err, app.TimeZeroErr.Error())
	})
}

func TestUpdateNote(t *testing.T) {
	title := "Title"
	body := "Note body text"
	newTitle := "Updated title"
	newBody := "Updated body"
	milliSecs := time.Now().UnixMilli()

	t.Run("should update note with no errors", func(t *testing.T) {
		assert := assert.New(t)
		note, _ := app.NewNote(title, body, milliSecs)
		err := note.UpdateNote(newTitle, newBody)
		assert.Nil(err)
	})

	t.Run("should update note title and body", func(t *testing.T) {
		assert := assert.New(t)
		note, _ := app.NewNote(title, body, milliSecs)
		note.UpdateNote(newTitle, newBody)
		assert.Equal(newTitle, note.Title())
		assert.Equal(newBody, note.Body())
	})

	t.Run("should return correct errors", func(t *testing.T) {
		assert := assert.New(t)
		note, _ := app.NewNote(title, body, milliSecs)
		err := note.UpdateNote("  ", newBody)
		assert.EqualError(err, app.EmptyTitleErr.Error())
		err = note.UpdateNote(newTitle, "  ")
		assert.EqualError(err, app.EmptyBodyErr.Error())
	})
}
