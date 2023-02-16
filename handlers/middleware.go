package handlers

import (
	"context"
	"net/http"

	"github.com/harshitasao/Microservices-Go/data"
)

// adding middleware means it is a httphandler and it allows to run chain of handlers for the request/response
// and the order of the chain depends on the code how you specify it

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Products{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			// data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}
		// add the product in the context - meaning that the data provided by the user is secure and in proper format
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// call the next handler that could be another middleware in the chain or could be the final handler
		next.ServeHTTP(rw, r)
	})
}
