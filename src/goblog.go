package main

import (
	"net/http"
	"html/template"
	"fmt"
	"log"
	"io/ioutil"
)

type Post struct {
	// TODO Serialized as a markdown file
	Title string
	Body []byte 
}

func titleToFilename(title string) string {
	return "posts/" + title + ".txt"
}

/*
func (p *Post) save() error {
	// TODO formatter
	filename := titleToFilename(p.Title) 
	return ioutil.WriteFile(filename, p.Body, 0600)
}
*/

func loadPost(title string) (*Post, error){
	filename := titleToFilename(title) 
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	} else {
		return &Post{Title: title, Body: body}, nil
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request){
	requestedPost := r.URL.Path[1:]
	post, err := loadPost(requestedPost)
	if err != nil {
		fmt.Fprintf(w, "%s does not exist yet!", requestedPost)
	} else {
		t, err := template.ParseFiles("templates/base.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, post)
	}
}

func main() {
	http.Handle("/static/", //final url can be anything
      http.StripPrefix("/static/",
         http.FileServer(http.Dir("static")))) 
	http.HandleFunc("/", mainHandler) 
   log.Fatal(http.ListenAndServe(":8080", nil))
}
