package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Routes        []Route
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Routes:        []Route{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Routes = append(s.Routes, Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)

	for _, route := range s.Routes {
		s.Router.Method(route.Method, route.Path, route.Handler)
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
