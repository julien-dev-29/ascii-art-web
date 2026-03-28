package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/ascii-art", handleAscii)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Ascii Art",
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	tmpl.Execute(w, data)
}

func handleAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result := GenerateAscii(text, banner)

	data := map[string]string{
		"Title":  "Ascii Art",
		"Result": result,
	}

	tmpl.Execute(w, data)
}
