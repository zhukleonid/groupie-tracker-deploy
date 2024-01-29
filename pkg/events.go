package pkg

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpl *template.Template
	err error
)

const artistAPI = "https://groupietrackers.herokuapp.com/api/artists"

// Функция обработчика домашней страницы
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound) // 404
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed) // 405
		return
	}

	bandinfo, err := GetBandInfo(artistAPI)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError) // 500
		return
	}

	tmpl_files := []string{"./web/templates/layout.html", "./web/templates/index.html"}
	var tpls *template.Template
	tpls, err = template.ParseFiles(tmpl_files...)
	if err != nil {
		log.Fatal(err)
	}

	err = tpls.ExecuteTemplate(w, "layout", bandinfo)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError) // 500
		return
	}
}

// Функция обработчика страницы с данными о группе и артистах
func BandHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	band, err := GetBandByID(id, bands)
	if err != nil {
		// Обработка ошибки, например, вывод на страницу ошибки или редирект
		ErrorHandler(w, http.StatusInternalServerError) // 500
		return
	}

	tmpl_files := []string{"./web/templates/layout.html", "./web/templates/band.html"}
	var tpls *template.Template
	tpls, err = template.ParseFiles(tmpl_files...)
	if err != nil {
		log.Fatal(err)
	}

	tpls.ExecuteTemplate(w, "layout", band)
}

// Функция обработчика ошибок
func ErrorHandler(w http.ResponseWriter, statusCode int) {
	
	w.WriteHeader(statusCode)
	
	data := struct {
		StatusMsg  string
		StatusCode int
	}{
		"Ooops. Error ",
		statusCode,
	}

	tmpl_files := []string{"./web/templates/layout.html", "./web/templates/error.html"}
	var tpls *template.Template
	tpls, err = template.ParseFiles(tmpl_files...)
	if err != nil {
		log.Fatal(err)
	}

	err := tpls.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
		return
	}
}
