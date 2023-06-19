package handelers

import (
	"html/template"
	"net/http"
)

func Hand() {
	fs := http.FileServer(http.Dir("./template"))
	http.Handle("/template/", http.StripPrefix("/template/", fs))

	http.HandleFunc("/", UrlCheck)
}

func UrlCheck(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		home(w, r)
	case "/artiste":
		artiste(w, r)
	case "/map":
		carte(w, r)
	default:

	}
}

func home(w http.ResponseWriter, r *http.Request) {

}

func artiste(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./template/HTML/artiste.html")
	if err != nil {

	}
	errr := tmp.ExecuteTemplate(w, "artiste", nil)
	if errr != nil {

	}
}

func carte(w http.ResponseWriter, r *http.Request) {

}
