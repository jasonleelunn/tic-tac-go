package server

import (
	"log"
	"net/http"
)

func Serve() {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", newGameHandler)

	router.HandleFunc("GET /status", ensureSession(statusHandler))

	router.HandleFunc("POST /cells/{cell}", ensureSession(cellHandler))

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":4040", router)

	if err != nil {
		log.Fatal("Server encountered an error: ", err)
	}
}
