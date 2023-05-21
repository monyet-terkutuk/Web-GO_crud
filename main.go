package main

import (
	"learn-web_crud/config"
	categories "learn-web_crud/controllers/categoriescontroller"
	home "learn-web_crud/controllers/homecontroller"
	product "learn-web_crud/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Halaman
	http.HandleFunc("/", home.Welcome)

	// 2. Products
	http.HandleFunc("/products", product.Index)

	// 3. Categories
	http.HandleFunc("/categories", categories.Index)
	http.HandleFunc("/categories/add", categories.Add)
	http.HandleFunc("/categories/edit", categories.Edit)
	http.HandleFunc("/categories/delete", categories.Delete)

	log.Println("Server Running")
	http.ListenAndServe(":8001", nil)
}
