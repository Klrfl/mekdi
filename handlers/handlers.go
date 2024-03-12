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
		Description: "",
	}

	var DBQuery string
	var rows *sql.Rows
	var err error

	if len(menuID) == 1 {
		DBQuery = "select * from menu where id = $1"
		rows, err = database.DB.Query("select * from menu where id = $1", menuID[0])
	} else {
		DBQuery = "select * from menu"
		rows, err = database.DB.Query(DBQuery)
	}

	if err != nil {
		return nil, fmt.Errorf("error when querying database: %s", err)
	}

	defer rows.Close()

	var description sql.NullString
	for rows.Next() {
		err := rows.Scan(
			&menu.Name,
			&menu.ServingSize,
			&menu.Ingredients,
			&menu.Tag,
			&menu.Allergy,
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
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			menu.Description = description.String
		}

		menuList = append(menuList, menu)
	}

	if err := rows.Err(); err == sql.ErrNoRows {
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
		log.Println(fmt.Errorf("error when getting menu: %s", err))
	}

	data := map[string]interface{}{
		"title": "Mekdi App",
		"data":  menuList,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	menuID, UUIDerror := uuid.Parse(r.PathValue("id"))

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
		"title": fmt.Sprintf("%s – Mekdi App", menuItem[0].Name),
		"data":  &menuItem[0],
	}

	tmpl.Execute(w, data)
}

func CreateNewMenu(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("error when parsing menu item data %s", err), http.StatusInternalServerError)
	}

	formData := r.Form
	menuID := uuid.New()
	menuName := formData.Get("menu-name")
	menuDescription := formData.Get("menu-description")
	menuServingSize := formData.Get("menu-serving-size")
	menuIngredients := formData.Get("menu-ingredients")
	menuTag := formData.Get("menu-tag")
	menuAllergy := formData.Get("menu-allergy")

	query := `insert into menu(id, name, description, serving_size, ingredients, tag, allergy)
  values($1, $2, $3, $4, $5, $6, $7)`

	_, err = database.DB.Exec(
		query,
		menuID,
		menuName,
		strings.TrimSpace(menuDescription),
		menuServingSize,
		menuIngredients,
		menuTag,
		menuAllergy,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("error when creating new menu item: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/menu/%s", menuID), http.StatusSeeOther)
}

func EditMenuItem(w http.ResponseWriter, r *http.Request) {
	IDParam := strings.TrimPrefix(r.URL.Path, "/menu/")
	menuID, UUIDerror := uuid.Parse(IDParam)

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

	query := `update menu set 
  name = $2,
  description = $3,
  serving_size = $4,
  ingredients = $5,
  tag = $6,
  allergy = $7
  where id = $1`

	_, err = database.DB.Exec(
		query,
		menuID,
		menuName,
		menuDescription,
		menuServingSize,
		menuIngredients,
		menuTag,
		menuAllergy,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("error when updating menu item: %s", err), http.StatusInternalServerError)
		return
	}

	successComponent := path.Join("views", "components", "success.html")
	tmpl := template.Must(template.ParseFiles(successComponent))
	tmpl.Execute(w, nil)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	menuID, UUIDerror := uuid.Parse(r.PathValue("id"))
	if UUIDerror != nil {
		http.Error(w, fmt.Sprintf("error when processing menu ID", UUIDerror), http.StatusBadRequest)
		return
	}

	query := "delete from menu where id = $1"
	_, err := database.DB.Exec(query, menuID)

	if err != nil {
		http.Error(w, fmt.Sprintf("error when deleting menu item: %s", err), http.StatusInternalServerError)
		return
	}

	deletedComponent := path.Join("views", "components", "deleted.html")
	tmpl := template.Must(template.ParseFiles(deletedComponent))
	tmpl.Execute(w, nil)
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
		var description sql.NullString

		for rows.Next() {
			err := rows.Scan(
				&menuItem.ID,
				&menuItem.Name,
				&menuItem.Tag,
				&menuItem.ServingSize,
				&menuItem.Ingredients,
				&menuItem.Description,
			)

			if err != nil {
				http.Error(w, fmt.Sprintf("error when getting menu items: %s", err), http.StatusInternalServerError)
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
