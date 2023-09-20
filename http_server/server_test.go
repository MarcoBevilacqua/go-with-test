package httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {

	t.Run("returns pepper's score", func (t *testing.T) {

		request, _ := http.NewRequest(http.methodGet, "players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		
		got := response.Body.string()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}