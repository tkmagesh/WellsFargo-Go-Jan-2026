package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

type HandlerFn func(http.ResponseWriter, *http.Request)

type AppServer struct {
	routes      map[string]HandlerFn
	middlewares []func(HandlerFn) HandlerFn
}

func (appServer *AppServer) AddHandler(path string, handlerFn HandlerFn) {
	for i := len(appServer.middlewares) - 1; i >= 0; i-- {
		handlerFn = appServer.middlewares[i](handlerFn)
	}
	appServer.routes[path] = handlerFn
}

func (appServer *AppServer) AddMiddleware(middleware func(HandlerFn) HandlerFn) {
	appServer.middlewares = append(appServer.middlewares, middleware)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes:      make(map[string]HandlerFn),
		middlewares: make([]func(HandlerFn) HandlerFn, 0),
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
	msg := GenerateGreet(r.Context())
	if r.Context().Err() == nil {
		fmt.Fprintln(w, msg)
	}
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

// Cross cutting concerns
func logWrapper(handlerFn HandlerFn) HandlerFn {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
		handlerFn(w, r)
	}
}

func timeoutWrapper(handlerFn HandlerFn) HandlerFn {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		handlerFn(w, r.WithContext(ctx))
		if ctx.Err() == context.DeadlineExceeded {
			http.Error(w, "request timedout", http.StatusRequestTimeout)
			return
		}
	}
}

// Business Logic
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
	appServer := NewAppServer()
	/*
		appServer.AddHandler("/", logWrapper(timeoutWrapper(IndexHandler)))
		appServer.AddHandler("/products", logWrapper(ProductsHandler))
		appServer.AddHandler("/customers", logWrapper(CustomersHandler))
	*/
	appServer.AddMiddleware(logWrapper)
	appServer.AddMiddleware(timeoutWrapper)
	appServer.AddHandler("/", IndexHandler)
	appServer.AddHandler("/products", ProductsHandler)
	appServer.AddHandler("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Fatalln(err)
	}
}
