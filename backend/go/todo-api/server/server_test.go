package server_test

import (
	"net/http"
	"testing"

	integration_test "github.com/clovisphere/todo/backend/go/todo-api/test/integration"
	"github.com/matryer/is"
)

func TestServerStart(t *testing.T) {
	integration_test.SkipIfShort(t)

	t.Run("starts the server and listens for requests", func(t *testing.T) {
		is := is.New(t)
		cleanup := integration_test.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:9000")
		is.NoErr(err)
		is.Equal(http.StatusNotFound, resp.StatusCode)
	})
}
