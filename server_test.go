package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	t.Run("return Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		res := httptest.NewRecorder()

		PlayerServer(res, req)
		got := res.Body.String()
		want := "10"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func buildServer(t *testing.T) {
	t.Helper()

}
