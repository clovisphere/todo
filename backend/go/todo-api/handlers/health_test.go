package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clovisphere/todo/backend/go/todo-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/matryer/is"
)

func TestHealth(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		is := is.New(t)

		mux := chi.NewMux()
		handlers.Health(mux)
		code, _, _ := makeGetRequest(mux, "/health")
		is.Equal(http.StatusOK, code)
	})
}

func makeGetRequest(handler http.Handler, target string) (int, http.Header, string) {
	req := httptest.NewRequest(http.MethodGet, target, nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	result := res.Result()
	bodyBytes, err := io.ReadAll(result.Body)

	if err != nil {
		panic(err)
	}
	return result.StatusCode, result.Header, string(bodyBytes)
}
