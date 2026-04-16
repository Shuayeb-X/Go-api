package main

import (
	"encoding/json"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Give me Valid JSON", 400)
		return
	}

	newProduct.Id = len(productlist) + 1
	productlist = append(productlist, newProduct)

	sendData(w, newProduct, 201)
}
