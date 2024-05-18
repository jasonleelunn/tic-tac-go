package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jasonleelunn/tic-tac-go/game"
)

func main() {
	var g *game.GameState

	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		g = game.New()
		game.Page("X's Turn").Render(r.Context(), w)
	})

	router.HandleFunc("GET /status", func(w http.ResponseWriter, r *http.Request) {
		if g.IsFinished() {
			w.Write([]byte(fmt.Sprintf("%s Wins!", g.GetWinningTokenString())))
			return
		}

		currToken := g.GetCurrentTokenString()

		w.Write([]byte(fmt.Sprintf("%s's Turn", currToken)))
	})

	router.HandleFunc("POST /cells/{cell}", func(w http.ResponseWriter, r *http.Request) {
		if g.IsFinished() {
			w.WriteHeader(http.StatusForbidden)
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
			w.WriteHeader(http.StatusForbidden)
			return
		}

		w.Header().Add("HX-Trigger", "status-changed")
		w.Write([]byte(g.GetCurrentTokenString()))

		g.ChangeCurrentToken()
	})

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":4040", router)

	if err != nil {
		log.Fatal("Server encountered an error: ", err)
	}
}
