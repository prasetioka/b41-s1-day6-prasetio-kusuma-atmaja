package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Selamat Datang Di Website Saya"))
	})

	route.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Tentang Saya"))
	})

	route.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pusat Bantuan"))
	})

	fmt.Println("Server Running on port 8000")
	http.ListenAndServe("localhost:8000", route)
}