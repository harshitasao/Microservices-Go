package handlers

import (
	"net/http"

	"github.com/harshitasao/Microservices-Go/product-api/data"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//		201: noContentResponse
//  404: errorResponse
//  422: errorValidation

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	// fetch data/product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Println("[DEBUG] updating record id", prod.ID)
	err := data.UpdateProduct(prod)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
