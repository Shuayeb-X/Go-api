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
	handleCors(w)
	handlepreflightReq(w, r)

	sendData(w, productlist, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlepreflightReq(w, r)
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Give me Valid JSON", 400)
		return

	}

	newProduct.Id = len(productlist) + 1
	productlist = append(productlist, newProduct)

	sendData(w, productlist, 201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handlepreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /products", http.HandlerFunc(createProduct))

	fmt.Println("Server running on port :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the Server", err)
	}
}

// inti function
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
