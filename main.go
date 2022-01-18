package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome, Descricao string
	Quantidade      int
	Valor           float64
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta RDP", Descricao: "Branca com estampa atras", Quantidade: 2, Valor: 95},
		{Nome: "Camiseta Megadeth", Descricao: "Nova coleção", Quantidade: 1, Valor: 150},
		{"Guitarra", "Flying V Branca", 1, 2500},
		{"Notebook", "Mac PRO", 1, 22500},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}
