package memory_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/vmkevv/netforenotes/app"
	"github.com/vmkevv/netforenotes/interface/api/openapi"
	"github.com/vmkevv/netforenotes/store/memory"
)

func TestNewNote(t *testing.T) {
	note1, _ := app.NewNote("First", "First note content", time.Now().UnixMilli())
	note2, _ := app.NewNote("Second", "Second note content", time.Now().UnixMilli())
	t.Run("should add notes", func(t *testing.T) {
		assert := assert.New(t)
		store := memory.NewNotesStore()
		store.SaveNote(note1)
		assert.Equal(1, store.GetLen())
		store.SaveNote(note2)
		assert.Equal(2, store.GetLen())
	})
}

func TestUpdateNote(t *testing.T) {
	note1, _ := app.NewNote("First", "First note content", time.Now().UnixMilli())
	t.Run("should update note title and body", func(t *testing.T) {
		assert := assert.New(t)
		newTitle := "new title"
		newBody := "new body"
		store := memory.NewNotesStore()
		store.SaveNote(note1)

		note1.UpdateNote(newTitle, note1.Body())
		store.UpdateNote(note1)
		storeNote1, _ := store.GetNote(note1.ID())
		assert.Equal(storeNote1.Title(), newTitle)
		assert.Equal(storeNote1.Body(), note1.Body())

		note1.UpdateNote(note1.Title(), newBody)
		store.UpdateNote(note1)
		storeNote1, _ = store.GetNote(note1.ID())
		assert.Equal(storeNote1.Body(), newBody)
		assert.Equal(storeNote1.Title(), note1.Title())
	})
}

func TestGetNote(t *testing.T) {
	note1, _ := app.NewNote("First", "First note content", time.Now().UnixMilli())
	note2, _ := app.NewNote("Second", "Second note content", time.Now().UnixMilli())
	t.Run("should return the correct note", func(t *testing.T) {
		assert := assert.New(t)
		store := memory.NewNotesStore()
		store.SaveNote(note1)
		store.SaveNote(note2)
		storeNote1, err := store.GetNote(note1.ID())
		assert.Nil(err)
		assert.EqualValues(note1, *storeNote1)
	})
	t.Run("should return not found error if note does not exists", func(t *testing.T) {
		assert := assert.New(t)
		store := memory.NewNotesStore()
		store.SaveNote(note1)
		_, err := store.GetNote(ulid.Make())
		assert.Error(err, memory.NotFoundErr.Error())
	})
}

func TestDeleteNote(t *testing.T) {
	note1, _ := app.NewNote("First", "First note content", time.Now().UnixMilli())
	note2, _ := app.NewNote("Second", "Second note content", time.Now().UnixMilli())
	t.Run("should delete notes correctly", func(t *testing.T) {
		assert := assert.New(t)
		store := memory.NewNotesStore()
		store.SaveNote(note1)
		store.SaveNote(note2)

		store.DeleteNote(note1.ID())
		assert.Equal(1, store.GetLen())

		storedNote2, _ := store.GetNote(note2.ID())
		assert.EqualValues(note2, *storedNote2)

		store.DeleteNote(note2.ID())
		assert.Equal(0, store.GetLen())
	})
}

func TestGetNotes(t *testing.T) {
	now := time.Now()
	note0, _ := app.NewNote("farst", "Farst note content", now.Add(-time.Second).UnixMilli())
	note1, _ := app.NewNote("First", "first note content", now.UnixMilli())
	note2, _ := app.NewNote("Second", "Second note content", now.Add(time.Second*3).UnixMilli())
	note3, _ := app.NewNote("Lorem", "Ipsum dor sit amet", now.Add(-time.Second*2).UnixMilli())
	note4, _ := app.NewNote("I like solid", "I think solid is the best front-end framework, way better than react ", now.Add(time.Second*1).UnixMilli())
	note5, _ := app.NewNote("Drink water", "Remeber always drink water, it's really important", now.Add(time.Hour*2).UnixMilli())
	store := memory.NewNotesStore()
	store.SaveNote(note0)
	store.SaveNote(note1)
	store.SaveNote(note2)
	store.SaveNote(note3)
	store.SaveNote(note4)
	store.SaveNote(note5)

	type testCase struct {
		desc     string
		params   openapi.GetNotesParams
		expected []app.Note
	}
	searchTerm1 := "rst"
	searchTerm2 := "EM"
	testCases := []testCase{
		{
			desc:     "Order asc by title and empty search",
			params:   openapi.GetNotesParams{Sort: openapi.Title, Order: openapi.Asc, Search: nil},
			expected: []app.Note{note5, note0, note1, note4, note3, note2},
		},
		{
			desc:     "Order desc by title and empty search",
			params:   openapi.GetNotesParams{Sort: openapi.Title, Order: openapi.Desc, Search: nil},
			expected: []app.Note{note2, note3, note4, note1, note0, note5},
		},
		{
			desc:     "Order asc by date and empty search",
			params:   openapi.GetNotesParams{Sort: openapi.Date, Order: openapi.Asc, Search: nil},
			expected: []app.Note{note3, note0, note1, note4, note2, note5},
		},
		{
			desc:     "Order desc by date and empty search",
			params:   openapi.GetNotesParams{Sort: openapi.Date, Order: openapi.Desc, Search: nil},
			expected: []app.Note{note5, note2, note4, note1, note0, note3},
		},
		{
			desc:     fmt.Sprintf("Order asc by title and search \"%s\"", searchTerm1),
			params:   openapi.GetNotesParams{Sort: openapi.Title, Order: openapi.Asc, Search: &searchTerm1},
			expected: []app.Note{note0, note1},
		},
		{
			desc:     fmt.Sprintf("Order asc by title and search \"%s\"", searchTerm2),
			params:   openapi.GetNotesParams{Sort: openapi.Title, Order: openapi.Asc, Search: &searchTerm2},
			expected: []app.Note{note5, note3},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			assert := assert.New(t)
			notes := store.GetNotes(test.params)
			assert.Equal(len(test.expected), len(notes))
			if len(notes) == len(test.expected) {
				for i, note := range notes {
					assert.Equal(test.expected[i], note, fmt.Sprintf("Note[%d] is not expected", i))
				}
			}
		})
	}
}
