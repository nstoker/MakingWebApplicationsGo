package controller

import (
	"net/http"
	"text/template"
)

var (
	homeController home
	shopController shop
)

// Startup startup the controller
func Startup(template map[string]*template.Template) {
	homeController.homeTemplate = template["home.html"]
	shopController.shopTemplate = template["shop.html"]

	homeController.registerRoutes()
	shopController.registerRoutes()

	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
