package router

import (
	"go-cutup/handler"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//our routes
	r.Get("/", handler.IndexHandler)
	r.Post("/message", handler.MessageHandler)

	//favicon
	r.Get("/favicon.ico", handler.FaviconHandler)

	//custom 404 and 405 handlers
	r.NotFound(handler.NotFoundHandler)
	r.MethodNotAllowed(handler.MethodNotAllowedHandler)

	//handle static files
	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	return r
}
