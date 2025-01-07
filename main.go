package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Artic struct {
	Title   string
	Content string
}

func main() {

	fmt.Println("hello world")

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmp1 := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Artic{
			"Articles": {
				{Title: "Article 1", Content: "Content of Article 1"},
				{Title: "Article 2", Content: "Content of Article 2"},
				{Title: "Article 3", Content: "Content of Article 3"},
			},
		}
		tmp1.Execute(w, films)
	}

	http.HandleFunc("/", handler1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
