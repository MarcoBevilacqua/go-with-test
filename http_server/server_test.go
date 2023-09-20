package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct{
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {	
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd": 10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("returns pepper's score", func (t *testing.T) {

		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
	
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func (t *testing.T) {
		request := newGetScoreRequest("Marco")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got:= response.Result().StatusCode
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d, want %d", got, want)
		}
		
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}

	server := &PlayerServer{&store}
	t.Run("it returns accepted on POST", func (t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})	

	t.Run("it records Win when POST", func (t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()
		
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d win calls, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("got %s, wanted %s", store.winCalls[0], player)
		}
	})
}

func newGetScoreRequest(playerName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", playerName), nil)
	return req
}

func newPostWinRequest(playerName string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", playerName), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}