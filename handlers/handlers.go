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
	"github.com/klrfl/mekdi/views"
)

func getMenu(menuID ...uuid.UUID) ([]models.Menu, error) {
	var menuList []models.Menu

	menu := models.Menu{
		ServingSize: "",
		Ingredients: "",
		Allergy:     "",
		Tag:         "",
		Description: "",
	}

	var servingSize sql.NullString
	var ingredients sql.NullString
	var tag sql.NullString
	var allergy sql.NullString
	var description sql.NullString

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
			&description,
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
		if description.Valid {
			menu.Description = description.String
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
			&description,
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
		if description.Valid {
			menu.Description = description.String
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

	page := path.Join("views", "index.html")
	tmpl := views.RenderPage(page)

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

func HandleMenu(w http.ResponseWriter, r *http.Request) {
	IDParam := strings.TrimPrefix(r.URL.Path, "/menu/")
	menuID, err := uuid.Parse(IDParam)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		menuItem, err := getMenu(menuID)
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}

		page := path.Join("views", "menuItem.html")
		tmpl := views.RenderPage(page)

		data := map[string]*models.Menu{
			"data": &menuItem[0],
		}

		tmpl.Execute(w, data)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "error when parsing menu item data", http.StatusInternalServerError)
		}

		formData := r.Form
		menuID := uuid.New()
		menuName := formData.Get("menu-name")
		menuServingSize := formData.Get("menu-serving-size")
		menuIngredients := formData.Get("menu-ingredients")
		menuTag := formData.Get("menu-tag")
		menuAllergy := formData.Get("menu-allergy")

		query := "insert into menu(id, name, serving_size, ingredients, tag, allergy) values($1, $2, $3, $4, $5, $6)"
		_, err = database.DB.Exec(query, menuID, menuName, menuServingSize, menuIngredients, menuTag, menuAllergy)

		if err != nil {
			log.Println(err)
			http.Error(w, "error when creating new menu item", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.MethodPatch:
		err = r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error processing form data: %s", err), http.StatusInternalServerError)
			return
		}

		formData := r.Form
		menuName := formData["menu-name"][0]
		menuServingSize := formData["menu-serving-size"][0]
		menuIngredients := formData["menu-ingredients"][0]
		menuTag := formData["menu-tag"][0]
		menuAllergy := formData["menu-allergy"][0]

		query := "update menu set name=$2, serving_size=$3, ingredients=$4, tag=$5, allergy=$6 where id=$1"
		_, err := database.DB.Exec(query, menuID, menuName, menuServingSize, menuIngredients, menuTag, menuAllergy)
		if err != nil {
			http.Error(w, "error when updating menu item", http.StatusInternalServerError)
			return
		}

		successComponent := path.Join("views", "components", "success.html")
		tmpl := template.Must(template.ParseFiles(successComponent))
		tmpl.Execute(w, nil)

	case http.MethodDelete:
		_, err = database.DB.Exec("delete from menu where id = $1", menuID)
		if err != nil {
			http.Error(w, "error when deleting menu item", http.StatusInternalServerError)
			return
		}

		deletedPage := path.Join("views", "components", "deleted.html")
		tmpl := template.Must(template.ParseFiles(deletedPage))
		tmpl.Execute(w, nil)

	default:
		http.Error(w, "idk man, method not allowed I guess", http.StatusInternalServerError)
	}
}

func ServeNewMenuPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		newMenuItemPage := path.Join("views", "newItem.html")
		tmpl := views.RenderPage(newMenuItemPage)
		tmpl.Execute(w, nil)
	}
}
