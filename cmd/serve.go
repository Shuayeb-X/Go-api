package cmd

import (
	"fmt"
	"net/http"

	"exammple.com/global_router"
	"exammple.com/handlers"
)

func Serve() {
	mux := http.NewServeMux()

	// Register handlers by path. Use HandleFunc to pass the function value.
	mux.HandleFunc("/gproducts", handlers.GetProducts)
	mux.HandleFunc("/cproducts", handlers.CreateProduct)

	fmt.Println("Server running on port :3000")

	router := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("Error starting the Server", err)
	}
}
