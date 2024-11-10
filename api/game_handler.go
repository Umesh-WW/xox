package api

import (
	"net/http"

	"github.com/Umesh-WW/xox/types"
)

func HandleCreateGame(w http.ResponseWriter, r *http.Request) {
	user := types.User{}
	_ = user
	game := types.Game{}
	_ = game
}

func HandleMakeMove(w http.ResponseWriter, r *http.Request) {

}
