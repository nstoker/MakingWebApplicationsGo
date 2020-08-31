package controller

import (
	"net/http"
	"text/template"

	"github.com/nstoker/MakingWebApplicationsGo/src/viewmodel"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}
