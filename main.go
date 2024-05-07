package main

import (
	"log"
	"net/http"
)

type Token int

type Game struct {
	currentToken Token
	grid         map[string]bool
}

const (
	naught Token = iota
	cross
)

var (
	tokens = map[Token]string{
		naught: "O",
		cross:  "X",
	}
)

func main() {
	game := Game{currentToken: naught, grid: make(map[string]bool)}

	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		Page().Render(r.Context(), w)
	})

	router.HandleFunc("POST /cells/{cell}", func(w http.ResponseWriter, r *http.Request) {
		cell := r.PathValue("cell")

		if game.grid[cell] == true {
			w.WriteHeader(http.StatusTeapot)
			return
		}

		game.grid[cell] = true

		w.Write([]byte(tokens[game.currentToken]))

		if game.currentToken == naught {
			game.currentToken = cross
		} else {
			game.currentToken = naught
		}
	})

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":4040", router)

	if err != nil {
		log.Fatal("Server encountered an error: ", err)
	}
}
