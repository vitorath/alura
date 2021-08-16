package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.FindAllProducts())
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.FindProductById(productId)
	temp.ExecuteTemplate(w, "EditProduct", product)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Fail to converte price", err.Error())
	}

	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		log.Println("Fail to converte quantity", err.Error())
	}

	product := models.Products{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Quantity:    quantity,
	}

	product.CreateNewProduct()

	http.Redirect(w, r, "/", 301)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println("Fail to converte id", err.Error())
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Fail to converte price", err.Error())
	}

	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		log.Println("Fail to converte quantity", err.Error())
	}

	product := models.Products{
		Id:          id,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Quantity:    quantity,
	}

	product.UpdateProduct()

	http.Redirect(w, r, "/", 301)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}
