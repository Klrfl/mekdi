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
	router.SetupRoutes()

	port := "localhost:8080"

	log.Println(fmt.Sprintf("server running at http://%s", port))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("failed to start server")
	}
}
