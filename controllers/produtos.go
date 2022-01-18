package controllers

import (
	"go-web-application/models"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func Criar(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Criar", nil)
}
