package api

import (
	"team-management/members/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Handlers struct {
	ErrorHandler *ErrorHandler
	CreateMember *CreateMember
	FilterMember *FilterMember
	GetMember    *GetMember
	UpdateMember *UpdateMember
	DeleteMember *DeleteMember
}

type Router struct {
	handlers *Handlers
}

func NewRouter(handlers *Handlers) *Router {
	return &Router{
		handlers: handlers,
	}
}

func (r *Router) Register(logger utils.Logger) *chi.Mux {
	router := chi.NewRouter()
	errorHandler := r.handlers.ErrorHandler

	router.Use(middleware.Recoverer)
	router.Use(cors.AllowAll().Handler)
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger, NoColor: true}))

	router.Route("/members", func(c chi.Router) {
		c.Post("/", errorHandler.Handle(r.handlers.CreateMember.Handle))
		c.Get("/", errorHandler.Handle(r.handlers.FilterMember.Handle))
		c.Get("/{id}", errorHandler.Handle(r.handlers.GetMember.Handle))
		c.Put("/{id}", errorHandler.Handle(r.handlers.UpdateMember.Handle))
		c.Delete("/{id}", errorHandler.Handle(r.handlers.DeleteMember.Handle))
	})

	return router
}
