package api

import (
	"encoding/json"
	"net/http"

	"github.com/Umesh-WW/xox/types"
)

var games []types.Game

// game type -> game ID
func HandleCreateGame(w http.ResponseWriter, r *http.Request) error {
	game := types.Game{
		Id:    len(games),
		Board: types.Iboard{{"", "", ""}, {"", "", ""}, {"", "", ""}},
		State: types.Play,
	}
	games = append(games, game)
	// Send the created game as JSON response
	// todo - create a common utility function
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game)
	return nil
}

func HandleMakeMove(w http.ResponseWriter, r *http.Request) error {
	return nil
}
