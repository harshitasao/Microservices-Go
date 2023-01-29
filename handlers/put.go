package handlers

import (
	"Microservice-Go/Microservices-Go/data"
	"net/http"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	// fetch data/product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Println("[DEBUG] updating record id", prod.ID)
	err := data.UpdateProduct(&prod)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
