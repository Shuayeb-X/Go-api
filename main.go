package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id          int
	Title       string
	Description string
	Price       string
	ImgUrl      string
}

var productlist []Product

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Let's Start")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Eren. See you in Paradise. Keep going")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	if r.Method == "OPTIONS"{
		w.WriteHeader(200)
		return
	}

	sendData(w, productlist, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	

	if r.Method == "OPTIONS"{
		w.WriteHeader(200)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Give me Valid JSON", 400)
		
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
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandleFunc(helloHandler))
	mux.HandleFunc("GET /about", http.HandleFunc(aboutHandler))
	mux.HandleFunc("GET /products",http. HandleFunc(getProducts))
	mux.HandleFunc("POST /newproducts",http. HandleFunccreateProduct))

	fmt.Println("Server running on port :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the Server", err)
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