package controllers

import (
	"go-web-application/models"
	"log"
	"net/http"
	"strconv"
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

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		valor := r.FormValue("valor")

		qtdeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		valorConvertidoParaFloat64, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			log.Println("Erro na conversão do valor:", err)
		}

		models.CriarNovoProduto(nome, descricao, qtdeConvertidoParaInt, valorConvertidoParaFloat64)
	}

	http.Redirect(w, r, "/", 301)
}
