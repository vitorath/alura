package models

import (
	"web/db"
)

type Products struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func (p *Products) CreateNewProduct() {
	db := db.ConnectDatabase()
	defer db.Close()

	insertDatabase, err := db.Prepare("insert into products(name, description, price, quantity) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insertDatabase.Exec(p.Name, p.Description, p.Price, p.Quantity)
}

func (p *Products) UpdateProduct() {
	db := db.ConnectDatabase()
	defer db.Close()

	updateDatabase, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateDatabase.Exec(p.Name, p.Description, p.Price, p.Quantity, p.Id)
}

func DeleteProduct(productId string) {
	db := db.ConnectDatabase()
	defer db.Close()

	deleteProduct, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}
	deleteProduct.Exec(productId)
}

func FindProductById(productId string) Products {
	db := db.ConnectDatabase()
	defer db.Close()

	findProduct, err := db.Query("select * from products where id = $1", productId)
	if err != nil {
		panic(err.Error())
	}

	for findProduct.Next() {
		product := Products{}
		err = findProduct.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
		return product
	}

	return Products{}
}

func FindAllProducts() []Products {
	db := db.ConnectDatabase()
	defer db.Close()

	selectAllProducts, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	products := []Products{}

	for selectAllProducts.Next() {
		product := Products{}
		err = selectAllProducts.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}

	return products
}
