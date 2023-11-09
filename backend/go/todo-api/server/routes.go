package server

import "github.com/clovisphere/todo/backend/go/todo-api/handlers"

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)
}
