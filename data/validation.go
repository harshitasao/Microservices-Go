package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// NOTE: validating the input is good for your security so that input is in your bound and if any mistakes are made easy to give error messages

// NOTE: for the fields whose values are not easy to predict/available fo rthat we can add custom validator

// Using go validator library for doing struct and field validation.
// this uses tags for validating

// this wraps all the validators field errors
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

type Validation struct {
	validate *validator.Validate
}

// creating a new validation of validation type
func NewValidation() *Validation {
	// creating a new validator
	validate := validator.New()
	// adding custom validator
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

// validate finc
func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)
	if len(errs) == 0 {
		return nil
	}
	var returnErrs []ValidationError
	for _, err := range errs {
		// casting field error into validationError and then appending it to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}

// custom validator func
func validateSKU(fl validator.FieldLevel) bool {
	// SKU must be in the format abc-abc-abc
	// for putting sku in that format gonna be using regexp
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	sku := re.FindAllString(fl.Field().String(), -1)

	if len(sku) == 1 {
		return true
	}

	return false
}
