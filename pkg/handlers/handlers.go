package handlers

import (
	"net/http"

	"github.com/eylemabz/go-course/pkg/render"
)

//Home is the home page andler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.html")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about-page.html")
}
