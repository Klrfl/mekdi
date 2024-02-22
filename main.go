package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, "anjing error cuk", http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"name": "siapa kek",
		"data": "blom ada",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
