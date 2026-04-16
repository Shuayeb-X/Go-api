package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imgUrl"`
}

var productlist []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	sendData(w, productlist, 200)
}

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

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PATCH,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func globalRouter(mux *http.ServeMux) http.HandlerFunc {
	handlerAllReq := func(w http.ResponseWriter, r *http.Request) {
		handleCors(w)

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		mux.ServeHTTP(w, r)
	}

	return handlerAllReq
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /gproducts", http.HandlerFunc(getProducts))
	mux.Handle("POST /cproducts", http.HandlerFunc(createProduct))

	fmt.Println("Server running on port :3000")

	router := globalRouter(mux)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("Error starting the Server", err)
	}
}

// init function
func init() {
	prd1 := Product{
		Id:          1,
		Title:       "Dragon",
		Description: "Dragon is red",
		Price:       "200 tk",
		ImgUrl:      "https://www.tastingtable.com/img/gallery/how-to-eat-dragon-fruit-for-the-uninitiated/intro-1682966430.jpg",
	}

	productlist = append(productlist, prd1)
}
