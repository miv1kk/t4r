package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tpl := template.Must(template.ParseFiles(wd + "\\app\\templates\\index.html"))

	tpl.Execute(w, nil)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("app/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}