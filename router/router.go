package router

import (
	"net/http"

	"github.com/klrfl/mekdi/handlers"
)

func SetupRoutes() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/menu/", handlers.HandleMenu)
	http.HandleFunc("/menu/new/", handlers.ServeNewMenuPage)
	http.HandleFunc("/search", handlers.HandleSearch)

}
