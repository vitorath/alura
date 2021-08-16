package routes

import (
	"net/http"
	"web/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/edit", controllers.EditProduct)

	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)

}
