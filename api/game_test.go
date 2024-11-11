package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Umesh-WW/xox/api"
	"github.com/Umesh-WW/xox/types"
)

func TestCreate(t *testing.T) {
	// Reset the games slice to ensure a clean state for each test

	// Create a new HTTP request to the /create-game endpoint
	req, err := http.NewRequest("POST", "/create-game", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Record the HTTP response using httptest
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.HandleCreateGame)

	// Serve the HTTP request
	handler.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("expected status %v; got %v", http.StatusCreated, status)
	}

	// Decode the response body into a Game struct
	var game types.Game
	if err := json.NewDecoder(recorder.Body).Decode(&game); err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	// Verify the game was created with expected properties
	if game.Id != 0 {
		t.Errorf("expected game ID 0; got %v", game.Id)
	}

	expectedBoard := types.Iboard{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
	if game.Board != expectedBoard {
		t.Errorf("expected board %v; got %v", expectedBoard, game.Board)
	}

	if game.State != types.Play {
		t.Errorf("expected state 'play'; got %v", game.State)
	}

	// Check that the game was added to the games slice
	// if len(games) != 1 {
	// 	t.Errorf("expected 1 game in games slice; got %v", len(games))
	// }
}
