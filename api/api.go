package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type apiServer struct {
	listenAdrr string
}

type ApiError struct {
	Error string `json:"error"`
}
type apiFunc func(w http.ResponseWriter, r *http.Request) error

func NewApiServer(listenAddress string) *apiServer {
	return &apiServer{
		listenAdrr: listenAddress,
	}
}
func MakeHttpApiHandler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			// create common utlity
			w.Header().Set("Content-Type", "Application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&ApiError{Error: err.Error()})
		}
	}
}

func (serv *apiServer) Run() {
	// to-do - create a sub server and  log
	http.HandleFunc("/v1/create-game", MakeHttpApiHandler(HandleCreateGame))
	http.HandleFunc("v1/move", MakeHttpApiHandler(HandleMakeMove))
	log.Println("server runing on port address ...", serv.listenAdrr)
	http.ListenAndServe(serv.listenAdrr, nil)
}
