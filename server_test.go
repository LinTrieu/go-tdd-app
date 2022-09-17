package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	t.Run("return Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder

		PlayerServer(request, response)

		got := response.Body.string()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
