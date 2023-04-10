package handlers

import (
	"net/http"

	"github.com/harshitasao/Microservices-Go/product-api/data"
)

// swagger:route GET /products products ListProducts
// Returns a list of products
// responses:
// 		200: productsResponse

// ListAll func to get/return all the products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prod := data.GetProducts()

	err := data.ToJSON(prod, rw)
	if err != nil {
		// log error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests for that user needs to provide ID
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return

	default:
		p.l.Println("[Error] fetching the product", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = data.ToJSON(prod, rw)
	if err != nil {
		p.l.Println("[ERROR] serializing", err)
	}

}
