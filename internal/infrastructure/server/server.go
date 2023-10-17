package server

import (
	"net/http"

	"email-service/internal/infrastructure/server/handlers/emails"

	"github.com/gorilla/mux"
)

// Context holds the dependencies for the server.
type Context struct {
	Email *emails.Handler
}

// api holds the router and implements the Server interface.
type api struct {
	router http.Handler
}

// Server is the interface that wraps the Router method.
type Server interface {
	Router() http.Handler
}

// Router returns the router of the server.
func (a *api) Router() http.Handler {
	return a.router
}

// NewRouter returns a new instance of the server.
func NewRouter(c *Context) Server {
	a := &api{}

	// Set up CORS middleware

	r := mux.NewRouter()
	//r.Use(middlewares.Recovery)
	r.HandleFunc("/send", c.Email.Handle).Methods(http.MethodPost)
	a.router = r

	return a
}
