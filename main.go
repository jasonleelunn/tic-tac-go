package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /foo", func(w http.ResponseWriter, r *http.Request) {
		component := Page()
		component.Render(r.Context(), w)
	})

	router.HandleFunc("POST /clicked-hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Clicked Me!"))
	})

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":4040", router)

	if err != nil {
		log.Fatal("Server encountered an error: ", err)
	}
}
