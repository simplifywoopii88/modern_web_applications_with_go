package handlers

import (
	"net/http"

	"github.com/simplifywoopii88/modern_web_applications_with_go/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.html")
}

