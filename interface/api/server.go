//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./types.config.yaml ./openapi/openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./server.config.yaml ./openapi/openapi.yaml
package api

import (
	"context"
	"net/http"

	"github.com/oklog/ulid/v2"
	"github.com/vmkevv/netforenotes/app"
	"github.com/vmkevv/netforenotes/interface/api/openapi"
)

type HTTPServer struct {
	repo app.Repository
}

var _ openapi.StrictServerInterface = (*HTTPServer)(nil)

func NewHTTPServer(app app.Repository) HTTPServer {
	return HTTPServer{app}
}

func (server HTTPServer) AddNote(
	ctx context.Context,
	req openapi.AddNoteRequestObject,
) (openapi.AddNoteResponseObject, error) {
	note, err := app.NewNote(
		req.Body.Title,
		req.Body.Body,
		req.Body.CreatedAt,
	)
	if err != nil {
		return openapi.AddNote4XXJSONResponse{
			Body:       openapi.ErrorMessage{Message: err.Error()},
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	server.repo.SaveNote(note)
	return openapi.AddNote201Response{}, nil
}

func (server HTTPServer) DeleteNote(
	ctx context.Context,
	req openapi.DeleteNoteRequestObject,
) (openapi.DeleteNoteResponseObject, error) {
	noteId, err := ulid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	server.repo.DeleteNote(noteId)
	return openapi.DeleteNote204Response{}, nil
}

func (server HTTPServer) UpdateNote(
	ctx context.Context,
	req openapi.UpdateNoteRequestObject,
) (openapi.UpdateNoteResponseObject, error) {
	noteId, err := ulid.Parse(req.Body.Id)
	if err != nil {
		return nil, err
	}
	note, err := server.repo.GetNote(noteId)
	if err != nil {
		return openapi.UpdateNote4XXJSONResponse{
			Body:       openapi.ErrorMessage{Message: err.Error()},
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	err = note.UpdateNote(req.Body.Title, req.Body.Body)
	if err != nil {
		return openapi.UpdateNote4XXJSONResponse{
			Body:       openapi.ErrorMessage{Message: err.Error()},
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	server.repo.UpdateNote(*note)
	return openapi.UpdateNote204Response{}, nil
}

func (server HTTPServer) GetNote(
	ctx context.Context,
	req openapi.GetNoteRequestObject,
) (openapi.GetNoteResponseObject, error) {
	noteId, err := ulid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	savedNote, err := server.repo.GetNote(noteId)
	if err != nil {
		return openapi.GetNote4XXJSONResponse{
			Body:       openapi.ErrorMessage{Message: err.Error()},
			StatusCode: http.StatusNotFound,
		}, nil
	}
	note := openapi.Note{
		Id:        savedNote.ID().String(),
		Title:     savedNote.Title(),
		Body:      savedNote.Body(),
		CreatedAt: savedNote.CreatedAt().UnixMilli(),
	}
	return openapi.GetNote200JSONResponse(note), nil
}

func (server HTTPServer) GetNotes(
	ctx context.Context,
	req openapi.GetNotesRequestObject,
) (openapi.GetNotesResponseObject, error) {
	savedNotes := server.repo.GetNotes(req.Params)
	notes := make(openapi.NoteList, len(savedNotes))
	for i, savedNote := range savedNotes {
		notes[i] = openapi.NoteListItem{
			Id:        savedNote.ID().String(),
			Title:     savedNote.Title(),
			CreatedAt: savedNote.CreatedAt().UnixMilli(),
		}
	}
	return openapi.GetNotes200JSONResponse(notes), nil
}
