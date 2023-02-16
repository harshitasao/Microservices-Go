package handlers

import (
	"net/http"

	"github.com/harshitasao/Microservices-Go/data"
)

// this is add new product to the productList
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch data/product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(&prod)
}
