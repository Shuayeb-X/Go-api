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
	if r.Method != http.MethodGet {
		http.Error(w, "Please give me GET request", 400)
		return
	}
	json.NewEncoder(w).Encode(productlist)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on port :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the Server", err)
	}
}

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
