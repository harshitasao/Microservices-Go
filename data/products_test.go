package data

import (
	"testing"
)

// very simple unit test for testing the validate func

func TestValidate(t *testing.T) {
	// created the product
	p := &Product{
		Name:  "harshita",
		Price: 1.00,
		SKU:   "abc-efg-fhg",
	}

	// doing validation
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
