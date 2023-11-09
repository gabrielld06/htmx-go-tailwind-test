package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Server started")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	homePage := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("./src/index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}

		templ.Execute(w, films)
	}

	addFilm := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		templ := template.Must(template.ParseFiles("./src/index.html"))
		templ.ExecuteTemplate(w, "film-list-item", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/add-film/", addFilm)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
