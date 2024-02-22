package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/klrfl/mekdi/database"
	"github.com/klrfl/mekdi/models"
)

func getMenus() ([]models.Menu, error) {
	rows, err := database.DB.Query("SELECT * FROM menu")
	if err != nil {
		return nil, fmt.Errorf("error when querying database")
	}

	defer rows.Close()

	var menuList []models.Menu

	var servingSize sql.NullString
	var ingredients sql.NullString
	var tag sql.NullString
	var allergy sql.NullString

	for rows.Next() {
		menu := models.Menu{
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, "error when rendering frontend!", http.StatusInternalServerError)
	}

	menuList, err := getMenus()
	if err != nil {
		log.Println(err)
	}

	data := map[string]interface{}{
		"name": "Ronald McDonald",
		"data": menuList,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
