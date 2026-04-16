package main

import (
	"net/http"
)

func globalRouter(mux *http.ServeMux) http.HandlerFunc {
	handlerAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PATCH,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		mux.ServeHTTP(w, r)
	}

	return handlerAllReq
}
