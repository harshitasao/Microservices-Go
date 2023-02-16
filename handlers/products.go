// Package classification of Product API
// Documentation of Product API

//  Schemes: http
//  BasePath: /
//  Version: v1

// 	Consumes:
// 	- application/json

// 	Produces:
// 	- application/json

// this gives go-swagger a chance to generate the docs
// swagger:meta

// swagger spec is used to generate co-gen in other language as well
// this API only deals with JSON

// top level documentaion helps in letting people know whats the intention of the service

package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/harshitasao/Microservices-Go/data"

	"github.com/gorilla/mux"
)

// struct for the documentation
// list of products return in response
// swagger: response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger: response noContent
type productNoContent struct {
}

// here we are defining in docs that this parameter is required for delete
// this relates to the func called deleteProduct
// swagger: parameters deleteProduct
type productIDParameterWrapper struct {
	// id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// keyProduct is the key of the product object
type KeyProduct struct{}

// :(
type Products struct {
	l *log.Logger
	v *data.Validation
}

// :(
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is the error message for defining that the product path is invalid
// so we just assigned the error message to a variable so that instead of using the whole message we can just use the variable
var ErrInvalidProductPath = fmt.Errorf("Invalid Path the path should be /products/[id]")

// generic error message returned by the server
type GenericError struct {
	Message string `json:"message"`
}

// collection of all the validation error messages
type ValidationError struct {
	Message string `json:"message"`
}

// getProductID return the product id from the URL
// it will panic if the id is not an integer but this will never happen as the router will make sure that it is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the product id to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}
	return id
}
