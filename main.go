package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
)

type me struct {
	Project []string
}

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("htmls/*"))
}
func main() {
	routes := httprouter.New()
	routes.GET("/", index)
	routes.GET("/project", projectShow)
	http.ListenAndServe(":8080", routes)
	appengine.Main()
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmp.ExecuteTemplate(w, "index.gohtml", nil)
}
func projectShow(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	k := me{
		Project: []string{
			"Go Go Go",
			"Prediction Goverment Project",
			"Golang_Backend",
		},
	}
	err := tmp.ExecuteTemplate(w, "projects.gohtml", k)
	if err != nil {
		log.Fatalln(err)
	}
}
