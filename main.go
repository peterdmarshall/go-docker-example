package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cockroachdb/pebble"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("Couldn't read PORT")
	}

	path, ok := os.LookupEnv("DB_PATH")
	if !ok {
		log.Fatal("Couldn't read DB_PATH")
	}

	_, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	log.Printf("Starting server on :%s", port)
	http.ListenAndServe(":"+port, r)
}