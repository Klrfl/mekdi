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

	http.HandleFunc("/menu/", handlers.HandleMenuByID)
	http.HandleFunc("/menu/new/", handlers.HandleNewMenuItem)
	http.HandleFunc("/menu/edit/", handlers.EditMenuByID)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
