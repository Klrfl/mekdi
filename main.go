package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/google/uuid"
	"github.com/klrfl/mekdi/database"
)

type Menu struct {
	ID          uuid.UUID
	Name        string
	ServingSize string
	Ingredients string
	Tag         string
	Allergy     string
	Energy      float64
	Protein     float64
	TotalFat    float64
	SatFat      float64
	TransFat    float64
	Chol        float64
	Carbs       float64
	TotalSugar  float64
	AddedSugar  float64
	Sodium      float64
	Description string
}

func GetMenus() ([]Menu, error) {
	rows, err := database.DB.Query("SELECT * FROM menu")
	if err != nil {
		return nil, fmt.Errorf("error when querying database")
	}

	defer rows.Close()

	var menuList []Menu

	var servingSize sql.NullString
	var ingredients sql.NullString
	var tag sql.NullString
	var allergy sql.NullString

	for rows.Next() {
		menu := Menu{
			ServingSize: "no serving size",
			Ingredients: "no ingredients",
			Allergy:     "no allergies",
			Tag:         "no tags",
		}
		if err := rows.Scan(
			&menu.Name,
			&servingSize,
			&ingredients,
			&tag,
			&allergy,
			&menu.Energy,
			&menu.Protein,
			&menu.TotalFat,
			&menu.SatFat,
			&menu.TransFat,
			&menu.Chol,
			&menu.Carbs,
			&menu.TotalSugar,
			&menu.AddedSugar,
			&menu.Sodium,
			&menu.Description,
			&menu.ID,
		); err != nil {
			return nil, err
		}

		if servingSize.Valid {
			menu.ServingSize = servingSize.String
		}
		if ingredients.Valid {
			menu.Ingredients = ingredients.String
		}
		if tag.Valid {
			menu.Tag = tag.String
		}
		if allergy.Valid {
			menu.Allergy = allergy.String
		}

		menuList = append(menuList, menu)
	}

	if err := rows.Err(); err != nil {
		return menuList, err
	}
	return menuList, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, "anjing error cuk", http.StatusInternalServerError)
	}

	menuList, err := GetMenus()
	if err != nil {
		log.Println(err)
	}
	data := map[string]interface{}{
		"name": "siapa kek",
		"data": menuList,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	database.Init()

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
