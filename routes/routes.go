package routes

import (
	"go-web-application/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/criar", controllers.Criar)
	http.HandleFunc("/insert", controllers.Insert)
}
