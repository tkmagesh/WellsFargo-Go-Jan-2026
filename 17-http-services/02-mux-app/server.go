package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

// Application specific logic
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	msg := GenerateGreet(r.Context())
	if r.Context().Err() == nil {
		fmt.Fprintln(w, msg)
	}
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(&products); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err == nil {
		products = append(products, newProduct)
		// http.Error(w, "product added", http.StatusCreated)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "new product created")
		return
	}
	http.Error(w, "request payload error", http.StatusBadRequest)
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The list of customers will be served")
}

func GenerateGreet(ctx context.Context) string {
	// simulate time consuming operation

	for range 5 {
		select {
		case <-ctx.Done():
			log.Println("[GenerateGreet] : request timedout")
			return ""
		default:
			fmt.Print(".")
			time.Sleep(time.Second)
		}
	}
	return "Hello, World!"
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/products", GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", AddNewProduct).Methods(http.MethodPost)
	router.HandleFunc("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln(err)
	}
}
