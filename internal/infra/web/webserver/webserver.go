package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Method string 
	Path string 
	Handler http.Handler
}

type WebServer struct {
	Router        chi.Router
	// Handlers      map[string]http.HandlerFunc
	Handlers      []Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		// Handlers:      make(map[string]http.HandlerFunc),
		Handlers:      []Handler{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	h := Handler{
		Method: method,
		Path: path,
		Handler: handler,
	}
	// s.Handlers[path] = handler
	s.Handlers = append(s.Handlers, h)
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		// s.Router.Handle(path, handler)
		s.Router.Method(handler.Method, handler.Path, handler.Handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
