package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/klrfl/mekdi/database"
	"github.com/klrfl/mekdi/models"
)

func getMenu(menuID ...uuid.UUID) ([]models.Menu, error) {
	var menuList []models.Menu

	menu := models.Menu{
		ServingSize: "no serving size",
		Ingredients: "no ingredients",
		Allergy:     "no allergies",
		Tag:         "no tags",
	}

	var servingSize sql.NullString
	var ingredients sql.NullString
	var tag sql.NullString
	var allergy sql.NullString

	if len(menuID) == 1 {
		row := database.DB.QueryRow("select * from menu where id = $1", menuID[0])

		if err := row.Scan(
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
		return menuList, nil
	}

	rows, err := database.DB.Query("select * from menu")
	if err != nil {
		return nil, fmt.Errorf("error when querying database")
	}

	defer rows.Close()

	for rows.Next() {
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
		return nil, err
	}

	return menuList, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, "error when rendering frontend!", http.StatusInternalServerError)
	}

	menuList, err := getMenu()
	if err != nil {
		log.Println(err)
	}

	data := map[string]interface{}{
		"data": menuList,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleMenuByID(w http.ResponseWriter, r *http.Request) {
	IDParam := strings.TrimPrefix(r.URL.Path, "/menu/")
	var filepath string

	menuID, err := uuid.Parse(IDParam)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	// TODO: get menu item by ID
	menuItem, err := getMenu(menuID)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return

	}
	filepath = path.Join("views", "menuItem.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, "error when rendering frontend!", http.StatusInternalServerError)
	}

	data := map[string]*models.Menu{
		"data": &menuItem[0],
	}

	tmpl.Execute(w, data)

}
