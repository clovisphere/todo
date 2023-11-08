package integration

import (
	"net/http"
	"testing"
	"time"

	"github.com/clovisphere/todo/backend/go/todo-api/server"
)

// CreateServer for testing on port 9000, returning a cleanup function that stops the server.
//
// Usage:
// cleanup := CreateServer()
// defer cleanup()
func CreateServer() func() {
	s := server.New(server.Options{
		Host: "localhost",
		Port: 9000,
	})

	go func() {
		if err := s.Start(); err != nil {
			panic(err)
		}
	}()

	for {
		_, err := http.Get("http://localhost:9000")
		if err == nil {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}

	return func() {
		if err := s.Stop(); err != nil {
			panic(err)
		}
	}
}

// SkipIfShort skips t if the "-short" flag is passed to "go test"
func SkipIfShort(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
}
