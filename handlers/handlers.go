package handlers

import (
	"Microservice-Go/Microservices-Go/data"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// // Hellohandler
// type Hello struct {
// 	l *log.Logger
// }

// func NewHello(l *log.Logger) *Hello {
// 	return &Hello{l}
// }

// // implementing that method on Hello instance(which are interface)
// func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	// NOTE:
// 	// avoid adding/using concret objects inside the handler so that we enough space to add the
// 	// objects according to the future cases(testability, future database use etc)

// 	h.l.Println("Hello, World!")
// 	// reading from the body and anything that implements the io.ReadCloser there we need to use
// 	// go library
// 	data, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(rw, "Oops", http.StatusBadRequest)
// 		return
// 	}
// 	// Using log we are getting data in the logs for the developer but not for the user
// 	// log.Printf("Data %s\n", data)

// 	// data for the user
// 	fmt.Fprintf(rw, "Hello %s\n", data)

// }

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface

// Commenting as now we are going to use much better and efficient approach
// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	// handle the request for a list of products
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		p.l.Println("PUT", r.URL.Path)
// 		// expect the id in the URI
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			p.l.Println("Invalid URI more than one id")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		if len(g[0]) != 2 {
// 			p.l.Println("Invalid URI more than one capture group")
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			p.l.Println("Invalid URI unable to convert to numer", idString)
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		p.updateProducts(id, rw, r)
// 		return
// 	}

// 	// catch all
// 	// if no method is satisfied return an error
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	data.AddProduct(prod)
}

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// adding middleware means it is a httphandler and it allows to run chain of handlers for the request/response
// and the order of the chain depends on the code how you specify it

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// took the old context from the rquest convereted it into the format we need and hence now placing that new context in the request
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)

	})
}
