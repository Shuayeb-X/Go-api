package main

import (
	"fmt"
	"net/http"
)

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
