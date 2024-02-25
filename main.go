package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/klrfl/mekdi/database"
	"github.com/klrfl/mekdi/handlers"
)

func main() {
	database.Init()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/menu/", handlers.HandleMenu)
	http.HandleFunc("/menu/new/", handlers.ServeNewMenuPage)

	port := "localhost:8080"

	log.Println(fmt.Sprintf("server running at http://%s", port))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("failed to start server")
	}
}
