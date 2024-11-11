package main

import (
	"flag"
	"net/http"

	"github.com/Umesh-WW/xox/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "todo")
	flag.Parse()
	http.HandleFunc("/v1/create-game", api.HandleCreateGame)
	http.HandleFunc("v1/move", api.HandleMakeMove)
	http.ListenAndServe(*listenAddr, nil)
}
