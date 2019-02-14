package hanlders

import (
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/didikprabowo/crud-go/conf"
)

var db = conf.DbConn()

type Post struct {
	ID    int
	Title string
	Body  string
}

var tmpl = template.Must(template.ParseGlob("templates/**/*.html"))

//index
func Pindex(w http.ResponseWriter, r *http.Request) {
	db := conf.DbConn()
	selDB, err := db.Query("SELECT * FROM posts ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Post{}
	res := []Post{}
	for selDB.Next() {
		var id int
		var title, body string
		err = selDB.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Title = strings.ToUpper(title)
		emp.Body = body[0:65]
		res = append(res, emp)
	}
	if err != nil {
		panic(err)
	}
	m := map[string]interface{}{
		"Results": res,
		"Titles":  "Blog Posts",
	}

	tmpl.ExecuteTemplate(w, "Index", m)

	defer db.Close()
}

//padd
func Padd(w http.ResponseWriter, r *http.Request) {
	db := conf.DbConn()
	if r.Method == "GET" {
		m := map[string]interface{}{
			"Titles": "Add Posts",
		}
		tmpl.ExecuteTemplate(w, "Add", m)
	} else {
		title := r.FormValue("title")
		body := r.FormValue("body")
		query, err := db.Prepare("INSERT INTO posts (title,body) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		query.Exec(title, body)
		defer db.Close()
		http.Redirect(w, r, "/index", 301)
	}
}

// Pedit...
func Pedit(w http.ResponseWriter, r *http.Request) {
	db := conf.DbConn()
	Nid := r.URL.Query().Get("id")
	query, err := db.Query("SELECT * FROM posts WHERE id=?", Nid)
	if err != nil {
		panic(err.Error())
	}
	pos := Post{}
	for query.Next() {
		var id int
		var title, body string
		err := query.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}
		pos.ID = id
		pos.Title = title
		pos.Body = body
		m := map[string]interface{}{
			"Titles":  "Edit Posts",
			"Results": pos,
		}
		tmpl.ExecuteTemplate(w, "EditPost", m)
		defer db.Close()
	}
}

// update
func Pupdate(w http.ResponseWriter, r *http.Request) {
	db := conf.DbConn()
	if r.Method == "POST" {
		id := r.FormValue("id")
		title := r.FormValue("title")
		body := r.FormValue("body")
		query, err := db.Prepare("UPDATE posts set title=?,body =? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		query.Exec(title, body, id)
		defer db.Close()
		http.Redirect(w, r, "/index", 301)
	}
}
func Pdelete(w http.ResponseWriter, r *http.Request) {
	db := conf.DbConn()
	Cid := r.URL.Query().Get("id")
	query, err := db.Prepare("DELETE FROM posts where id =?")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(Cid)
	defer db.Close()
	http.Redirect(w, r, "/index", 301)
}
