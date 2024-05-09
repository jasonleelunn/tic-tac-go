package main

import (
	"log"
	"net/http"

	"github.com/jasonleelunn/tic-tac-go/game"
)

func main() {
	state := game.New()

	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		game.Page().Render(r.Context(), w)
	})

	router.HandleFunc("POST /cells/{cell}", func(w http.ResponseWriter, r *http.Request) {
		cell := r.PathValue("cell")

		// disallow choosing a cell more than once
		if state.GetGridCell(cell) == true {
			w.WriteHeader(http.StatusTeapot)
			return
		}

		state.MarkGridCell(cell)

		w.Write([]byte(state.GetCurrentTokenString()))

		state.ChangeCurrentToken()
	})

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":4040", router)

	if err != nil {
		log.Fatal("Server encountered an error: ", err)
	}
}
