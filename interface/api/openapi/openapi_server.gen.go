// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /note)
	AddNote(ctx echo.Context) error

	// (PUT /note)
	UpdateNote(ctx echo.Context) error

	// (DELETE /note/{id})
	DeleteNote(ctx echo.Context, id string) error

	// (GET /note/{id})
	GetNote(ctx echo.Context, id string) error

	// (GET /notes)
	GetNotes(ctx echo.Context, params GetNotesParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AddNote converts echo context to params.
func (w *ServerInterfaceWrapper) AddNote(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddNote(ctx)
	return err
}

// UpdateNote converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateNote(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateNote(ctx)
	return err
}

// DeleteNote converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteNote(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteNote(ctx, id)
	return err
}

// GetNote converts echo context to params.
func (w *ServerInterfaceWrapper) GetNote(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNote(ctx, id)
	return err
}

// GetNotes converts echo context to params.
func (w *ServerInterfaceWrapper) GetNotes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNotesParams
	// ------------- Required query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, true, "sort", ctx.QueryParams(), &params.Sort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort: %s", err))
	}

	// ------------- Required query parameter "order" -------------

	err = runtime.BindQueryParameter("form", true, true, "order", ctx.QueryParams(), &params.Order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order: %s", err))
	}

	// ------------- Optional query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, false, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNotes(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/note", wrapper.AddNote)
	router.PUT(baseURL+"/note", wrapper.UpdateNote)
	router.DELETE(baseURL+"/note/:id", wrapper.DeleteNote)
	router.GET(baseURL+"/note/:id", wrapper.GetNote)
	router.GET(baseURL+"/notes", wrapper.GetNotes)

}

type AddNoteRequestObject struct {
	Body *AddNoteJSONRequestBody
}

type AddNoteResponseObject interface {
	VisitAddNoteResponse(w http.ResponseWriter) error
}

type AddNote201Response struct {
}

func (response AddNote201Response) VisitAddNoteResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type AddNote4XXJSONResponse struct {
	Body       ErrorMessage
	StatusCode int
}

func (response AddNote4XXJSONResponse) VisitAddNoteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type UpdateNoteRequestObject struct {
	Body *UpdateNoteJSONRequestBody
}

type UpdateNoteResponseObject interface {
	VisitUpdateNoteResponse(w http.ResponseWriter) error
}

type UpdateNote204Response struct {
}

func (response UpdateNote204Response) VisitUpdateNoteResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type UpdateNote4XXJSONResponse struct {
	Body       ErrorMessage
	StatusCode int
}

func (response UpdateNote4XXJSONResponse) VisitUpdateNoteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type DeleteNoteRequestObject struct {
	Id string `json:"id"`
}

type DeleteNoteResponseObject interface {
	VisitDeleteNoteResponse(w http.ResponseWriter) error
}

type DeleteNote204Response struct {
}

func (response DeleteNote204Response) VisitDeleteNoteResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteNote4XXJSONResponse struct {
	Body       ErrorMessage
	StatusCode int
}

func (response DeleteNote4XXJSONResponse) VisitDeleteNoteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetNoteRequestObject struct {
	Id string `json:"id"`
}

type GetNoteResponseObject interface {
	VisitGetNoteResponse(w http.ResponseWriter) error
}

type GetNote200JSONResponse Note

func (response GetNote200JSONResponse) VisitGetNoteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetNote4XXJSONResponse struct {
	Body       ErrorMessage
	StatusCode int
}

func (response GetNote4XXJSONResponse) VisitGetNoteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetNotesRequestObject struct {
	Params GetNotesParams
}

type GetNotesResponseObject interface {
	VisitGetNotesResponse(w http.ResponseWriter) error
}

type GetNotes200JSONResponse NoteList

func (response GetNotes200JSONResponse) VisitGetNotesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /note)
	AddNote(ctx context.Context, request AddNoteRequestObject) (AddNoteResponseObject, error)

	// (PUT /note)
	UpdateNote(ctx context.Context, request UpdateNoteRequestObject) (UpdateNoteResponseObject, error)

	// (DELETE /note/{id})
	DeleteNote(ctx context.Context, request DeleteNoteRequestObject) (DeleteNoteResponseObject, error)

	// (GET /note/{id})
	GetNote(ctx context.Context, request GetNoteRequestObject) (GetNoteResponseObject, error)

	// (GET /notes)
	GetNotes(ctx context.Context, request GetNotesRequestObject) (GetNotesResponseObject, error)
}

type StrictHandlerFunc func(ctx echo.Context, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// AddNote operation middleware
func (sh *strictHandler) AddNote(ctx echo.Context) error {
	var request AddNoteRequestObject

	var body AddNoteJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AddNote(ctx.Request().Context(), request.(AddNoteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AddNote")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AddNoteResponseObject); ok {
		return validResponse.VisitAddNoteResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// UpdateNote operation middleware
func (sh *strictHandler) UpdateNote(ctx echo.Context) error {
	var request UpdateNoteRequestObject

	var body UpdateNoteJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateNote(ctx.Request().Context(), request.(UpdateNoteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateNote")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(UpdateNoteResponseObject); ok {
		return validResponse.VisitUpdateNoteResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// DeleteNote operation middleware
func (sh *strictHandler) DeleteNote(ctx echo.Context, id string) error {
	var request DeleteNoteRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteNote(ctx.Request().Context(), request.(DeleteNoteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteNote")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteNoteResponseObject); ok {
		return validResponse.VisitDeleteNoteResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// GetNote operation middleware
func (sh *strictHandler) GetNote(ctx echo.Context, id string) error {
	var request GetNoteRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetNote(ctx.Request().Context(), request.(GetNoteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetNote")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetNoteResponseObject); ok {
		return validResponse.VisitGetNoteResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// GetNotes operation middleware
func (sh *strictHandler) GetNotes(ctx echo.Context, params GetNotesParams) error {
	var request GetNotesRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetNotes(ctx.Request().Context(), request.(GetNotesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetNotes")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetNotesResponseObject); ok {
		return validResponse.VisitGetNotesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}
