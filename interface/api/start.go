package api

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vmkevv/netforenotes/interface/api/openapi"
	"github.com/vmkevv/netforenotes/store/memory"
)

func StartHTTPServer() {
	port := os.Getenv("APP_PORT")
	repo := memory.NewNotesStore()
	server := NewHTTPServer(repo)
	handler := openapi.NewStrictHandler(server, []openapi.StrictMiddlewareFunc{})
	app := echo.New()
	app.Use(middleware.CORS())
	openapi.RegisterHandlers(app, handler)
	app.Logger.Fatal(app.Start(":" + port))
}
