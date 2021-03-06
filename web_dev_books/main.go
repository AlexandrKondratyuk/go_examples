package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Name string
	DBStatus bool
}

func main() {
	p := Page{Name: "Alex"}

	templates := template.Must(template.ParseFiles("./templates/index.html"))

	db, _ := sql.Open("sqlite3", "dev.db")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		p.DBStatus = db.Ping() == nil

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		db.Close()
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
