package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jasonleelunn/tic-tac-go/game"
)

func main() {
	g := game.New()

	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		game.Page().Render(r.Context(), w)
	})

	router.HandleFunc("POST /cells/{cell}", func(w http.ResponseWriter, r *http.Request) {
		// do nothing if game is already over
		if g.IsFinished() {
			w.WriteHeader(http.StatusTeapot)
			return
		}

		cellNumber, err := strconv.Atoi(r.PathValue("cell"))

		if err != nil || cellNumber < 0 || cellNumber > 8 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = g.MarkGridCell(cellNumber)

		// disallow choosing a cell more than once
		if err != nil {
			w.WriteHeader(http.StatusTeapot)
			return
		}

		w.Write([]byte(g.GetCurrentTokenString()))

		g.ChangeCurrentToken()
	})

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":4040", router)

	if err != nil {
		log.Fatal("Server encountered an error: ", err)
	}
}
