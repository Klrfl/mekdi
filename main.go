package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/klrfl/mekdi/database"
	"github.com/klrfl/mekdi/router"
)

func main() {
	database.Init()
	mux := router.SetupRoutes()

	port := "localhost:8080"

	log.Println(fmt.Sprintf("server running at http://%s", port))
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to start server: %s", err))
	}
}
