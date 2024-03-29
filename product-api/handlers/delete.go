package handlers

import (
	"net/http"

	"github.com/harshitasao/Microservices-Go/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProducts
// Returns a list of products
// responses:
// 		201: noContentResponse
//  404: errorResponse
//  501: errorResponse

func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("[DEBUG] delete particular product", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)

}
