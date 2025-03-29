package main

import (
	"net/http"
	"text/template"

	"github.com/Toluxpersia/tolu.com/pkg"
)

func main() {
	pkg.Connect()
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
