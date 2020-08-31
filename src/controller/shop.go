package controller

import (
	"net/http"
	"text/template"

	"github.com/nstoker/MakingWebApplicationsGo/src/viewmodel"
)

type shop struct {
	shopTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleshop)
}

func (s shop) handleshop(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewShop()
	s.shopTemplate.Execute(w, vm)
}
