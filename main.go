package main

import (
	"github.com/klrfl/mekdi/database"
	"github.com/klrfl/mekdi/handlers"
	"log"
	"net/http"
)

func main() {
	database.Init()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/menu/", handlers.HandleMenu)
	http.HandleFunc("/menu/new/", handlers.ServeNewMenuPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
