package hanlders

import (
	"html/template"
	"net/http"
)

type ogps struct {
	Url string
}

type atribut struct {
	Title   string
	Keyword []string
	Ogp     []ogps
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/**/*.html"))
}

func (data atribut) Process() int {
	return 7
}

// func Hindex . . .
func Hindex(w http.ResponseWriter, r *http.Request) {

	data := atribut{
		Title:   "Home page",
		Keyword: []string{"Belajar", "Bermain"},
		Ogp:     []ogps{{Url: "http"}},
	}
	err := tpl.ExecuteTemplate(w, "indexing.html", data)
	if err != nil {
		panic(err)
	}
}
