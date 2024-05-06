package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		component := hello("Jason")
		component.Render(r.Context(), w)
	})

	router.HandleFunc("GET /foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got foo")
	})

	err := http.ListenAndServe(":4040", router)

	if err != nil {
		log.Fatal("Server encountered an error: ", err)
	}
}
