package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/partials/nav.html",
		"./ui/html/base.html",
		"./ui/html/home.html",
	}

	w.Header().Add("Server", "Go")
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal servererror", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying a specifc snippet with id %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
