package handlers

import (
	"net/http"

	"exammple.com/product"
	"exammple.com/utility"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utility.SendData(w, product.Productlist, 200)
}
