package routes

import (
	"GoCRUD/controllers"
	"net/http"
)

// CarregarRotas é utilizado para carregar as rotas do controller
func CarregarRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
