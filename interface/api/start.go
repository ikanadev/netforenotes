package api

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/vmkevv/netforenotes/interface/api/openapi"
	"github.com/vmkevv/netforenotes/store/memory"
)

func StartHTTPServer() {
	port := os.Getenv("APP_PORT")
	repo := memory.NewNotesStore()
	server := NewHTTPServer(repo)
	handler := openapi.NewStrictHandler(server, []openapi.StrictMiddlewareFunc{})
	app := echo.New()
	openapi.RegisterHandlers(app, handler)
	app.Logger.Fatal(app.Start(":" + port))
}
