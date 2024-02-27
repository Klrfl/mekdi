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

	query := "select * from menu"
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("error when querying database: %s", err))
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
		tmpl := views.Render404()
		tmpl.Execute(w, nil)
		return
	}

	page := path.Join("views", "index.html")
	tmpl := views.RenderPage(page)

	menuList, err := getMenu()
	if err != nil {
		log.Println(fmt.Sprintf("error when getting menu: %s", err))
	}

	data := map[string]interface{}{
		"title": "McD App",
		"data":  menuList,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleMenu(w http.ResponseWriter, r *http.Request) {
	IDParam := strings.TrimPrefix(r.URL.Path, "/menu/")
	menuID, UUIDerror := uuid.Parse(IDParam)

	switch r.Method {
	case http.MethodGet:
		if UUIDerror != nil {
			tmpl := views.Render404()
			tmpl.Execute(w, nil)
		}

		menuItem, err := getMenu(menuID)
		if err == sql.ErrNoRows {
			tmpl := views.Render404()
			tmpl.Execute(w, nil)
			return
		}

		page := path.Join("views", "menuItem.html")
		tmpl := views.RenderPage(page)

		data := map[string]interface{}{
			"title": fmt.Sprintf("%s – Mcd App", menuItem[0].Name),
			"data":  &menuItem[0],
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
		menuDescription := formData.Get("menu-description")
		menuServingSize := formData.Get("menu-serving-size")
		menuIngredients := formData.Get("menu-ingredients")
		menuTag := formData.Get("menu-tag")
		menuAllergy := formData.Get("menu-allergy")

		query := "insert into menu(id, name, description, serving_size, ingredients, tag, allergy) values($1, $2, $3, $4, $5, $6, $7)"
		_, err = database.DB.Exec(query, menuID, menuName, menuDescription, menuServingSize, menuIngredients, menuTag, menuAllergy)

		if err != nil {
			http.Error(w, fmt.Sprintf("error when creating new menu item: %s", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.MethodPatch:
		if UUIDerror != nil {
			tmpl := views.Render404()
			tmpl.Execute(w, nil)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error processing form data: %s", err), http.StatusInternalServerError)
			return
		}

		formData := r.Form
		menuName := formData.Get("menu-name")
		menuDescription := formData.Get("menu-description")
		menuServingSize := formData.Get("menu-serving-size")
		menuIngredients := formData.Get("menu-ingredients")
		menuTag := formData.Get("menu-tag")
		menuAllergy := formData.Get("menu-allergy")

		query := "update menu set name=$2, description=$3, serving_size=$4, ingredients=$5, tag=$6, allergy=$7 where id=$1"
		_, err = database.DB.Exec(query, menuID, menuName, menuDescription, menuServingSize, menuIngredients, menuTag, menuAllergy)

		if err != nil {
			http.Error(w, fmt.Sprintf("error when updating menu item: %s", err), http.StatusInternalServerError)
			return
		}

		successComponent := path.Join("views", "components", "success.html")
		tmpl := template.Must(template.ParseFiles(successComponent))
		tmpl.Execute(w, nil)

	case http.MethodDelete:
		query := "delete from menu where id = $1"
		_, err := database.DB.Exec(query, menuID)

		if err != nil {
			http.Error(w, fmt.Sprintf("error when deleting menu item: %s", err), http.StatusInternalServerError)
			return
		}

		deletedComponent := path.Join("views", "components", "deleted.html")
		tmpl := template.Must(template.ParseFiles(deletedComponent))
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

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	if query != "" {
		var menuList []models.Menu
		DBQuery := "select id, name, tag, serving_size, ingredients, description from menu where name ILIKE $1"
		rows, err := database.DB.Query(DBQuery, fmt.Sprintf("%%%s%%", query))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var menuItem models.Menu
		var servingSize sql.NullString
		var ingredients sql.NullString
		var tag sql.NullString
		var description sql.NullString

		for rows.Next() {
			err := rows.Scan(
				&menuItem.ID,
				&menuItem.Name,
				&tag,
				&servingSize,
				&ingredients,
				&description,
			)

			if err != nil {
				http.Error(w, fmt.Sprintf("error when getting menu items: %s", err.Error()), http.StatusInternalServerError)
			}

			if servingSize.Valid {
				menuItem.ServingSize = servingSize.String
			}
			if ingredients.Valid {
				menuItem.Ingredients = ingredients.String
			}
			if tag.Valid {
				menuItem.Tag = tag.String
			}
			if description.Valid {
				menuItem.Description = description.String
			}
			menuList = append(menuList, menuItem)
		}

		data := map[string][]models.Menu{
			"data": menuList,
		}

		searchResultsComponent := path.Join("views", "components", "searchResults.html")
		tmpl := template.Must(template.ParseFiles(searchResultsComponent))
		tmpl.Execute(w, data)
		return
	}

	data := map[string]interface{}{
		"title": "Search",
	}

	searchPage := path.Join("views", "search.html")
	tmpl := views.RenderPage(searchPage)
	tmpl.Execute(w, data)
}
