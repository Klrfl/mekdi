package router

import (
	"net/http"

	"github.com/klrfl/mekdi/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("GET /", handlers.IndexHandler)

	mux.HandleFunc("GET /menu/{id}", handlers.GetMenu)
	mux.HandleFunc("DELETE /menu/{id}", handlers.DeleteMenu)
	mux.HandleFunc("POST /menu/", handlers.CreateNewMenu)
	mux.HandleFunc("PATCH /menu/", handlers.EditMenuItem)

	mux.HandleFunc("GET /menu/new/", handlers.ServeNewMenuPage)
	mux.HandleFunc("GET /search/", handlers.HandleSearch)
	return mux
}
