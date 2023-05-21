// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/harshitasao/Microservices-Go/product-api/client/models"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProductCreated creates a UpdateProductCreated with default headers values
func NewUpdateProductCreated() *UpdateProductCreated {
	return &UpdateProductCreated{}
}

/*
UpdateProductCreated describes a response with status code 201, with default header values.

No content is returned by this API endpoint
*/
type UpdateProductCreated struct {
}

// IsSuccess returns true when this update product created response has a 2xx status code
func (o *UpdateProductCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update product created response has a 3xx status code
func (o *UpdateProductCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update product created response has a 4xx status code
func (o *UpdateProductCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update product created response has a 5xx status code
func (o *UpdateProductCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update product created response a status code equal to that given
func (o *UpdateProductCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UpdateProductCreated) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductCreated ", 201)
}

func (o *UpdateProductCreated) String() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductCreated ", 201)
}

func (o *UpdateProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductNotFound creates a UpdateProductNotFound with default headers values
func NewUpdateProductNotFound() *UpdateProductNotFound {
	return &UpdateProductNotFound{}
}

/*
UpdateProductNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type UpdateProductNotFound struct {
	Payload *models.GenericError
}

// IsSuccess returns true when this update product not found response has a 2xx status code
func (o *UpdateProductNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update product not found response has a 3xx status code
func (o *UpdateProductNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update product not found response has a 4xx status code
func (o *UpdateProductNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update product not found response has a 5xx status code
func (o *UpdateProductNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update product not found response a status code equal to that given
func (o *UpdateProductNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateProductNotFound) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProductNotFound) String() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductUnprocessableEntity creates a UpdateProductUnprocessableEntity with default headers values
func NewUpdateProductUnprocessableEntity() *UpdateProductUnprocessableEntity {
	return &UpdateProductUnprocessableEntity{}
}

/*
UpdateProductUnprocessableEntity describes a response with status code 422, with default header values.

Validation errors defined as an array of strings
*/
type UpdateProductUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this update product unprocessable entity response has a 2xx status code
func (o *UpdateProductUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update product unprocessable entity response has a 3xx status code
func (o *UpdateProductUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update product unprocessable entity response has a 4xx status code
func (o *UpdateProductUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update product unprocessable entity response has a 5xx status code
func (o *UpdateProductUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update product unprocessable entity response a status code equal to that given
func (o *UpdateProductUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *UpdateProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateProductUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateProductUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
