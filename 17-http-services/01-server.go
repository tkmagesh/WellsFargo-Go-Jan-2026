package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Cost    float64 `json:"cost"`
	InStock bool    `json:"-"`
}

// in-memory products store
var products []Product = []Product{
	{101, "Pen", 10, true},
	{102, "Pencil", 5, true},
	{103, "Marker", 50, false},
	{104, "Mouse", 400, true},
}

type AppServer struct {
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))

	// log the request
	fmt.Printf("%s - %s\n", r.Method, r.URL.Path)

	// serve the request
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello, World!")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(&products); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err == nil {
				products = append(products, newProduct)
				// http.Error(w, "product added", http.StatusCreated)
				w.WriteHeader(http.StatusCreated)
				fmt.Fprintln(w, "new product created")
				return
			}
			http.Error(w, "request payload error", http.StatusBadRequest)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "The list of customers will be served")
	default:
		http.Error(w, "Resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Fatalln(err)
	}
}
