package handlers

import (
	"encoding/json"
	"net/http"

	"exammple.com/product"
	"exammple.com/utility"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct product.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Give me Valid JSON", http.StatusBadRequest)
		return
	}

	newProduct.Id = len(product.Productlist) + 1
	product.Productlist = append(product.Productlist, newProduct)

	utility.SendData(w, newProduct, http.StatusCreated)
}
