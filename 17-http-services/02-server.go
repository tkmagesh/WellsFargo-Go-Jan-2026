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
	routes map[string]func(w http.ResponseWriter, r *http.Request)
}

func (appServer *AppServer) AddHandler(path string, handlerFn func(w http.ResponseWriter, r *http.Request)) {
	appServer.routes[path] = handlerFn
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	}
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlerFn := appServer.routes[r.URL.Path]; handlerFn != nil {
		handlerFn(w, r)
		return
	}
	http.Error(w, "Resource not found", http.StatusNotFound)
}

// Application specific logic
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The list of customers will be served")
}

/*
// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))

	// log the request
	fmt.Printf("%s - %s\n", r.Method, r.URL.Path)

	// serve the request
	switch r.URL.Path {
	case "/":

	case "/products":

	case "/customers":

	default:
		http.Error(w, "Resource not found", http.StatusNotFound)
	}

} */

func main() {
	appServer := NewAppServer()
	appServer.AddHandler("/", IndexHandler)
	appServer.AddHandler("/products", ProductsHandler)
	appServer.AddHandler("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Fatalln(err)
	}
}
