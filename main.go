package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/error.html"))

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/ascii-art", handleAscii)
	mux.HandleFunc("/download", handleDownload)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8000", mux)
}

func renderError(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	data := map[string]string{
		"Title":   strconv.Itoa(status) + " " + http.StatusText(status),
		"Message": message,
	}
	if err := tmpl.ExecuteTemplate(w, "error.html", data); err != nil {
		http.Error(w, message, status)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		renderError(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	data := map[string]string{
		"Title": "Ascii Art",
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		renderError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderError(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		renderError(w, "Text is required", http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		renderError(w, "Invalid banner", http.StatusBadRequest)
		return
	}
	result, err := GenerateAscii(text, banner)
	if err != nil {
		renderError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := map[string]string{
		"Title":  "Ascii Art",
		"Text":   text,
		"Banner": banner,
		"Result": result,
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		renderError(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderError(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		renderError(w, "Text is required", http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		renderError(w, "Invalid banner", http.StatusBadRequest)
		return
	}
	result, err := GenerateAscii(text, banner)
	if err != nil {
		renderError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename=\"ascii-art.txt\"")
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
