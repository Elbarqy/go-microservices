// Package classification of product API
//
// Documentation for Product API
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

package data

import "testing"

func TestCheckValidationFunctionality(t *testing.T) {
	p := &Product{
		Name:  "woo",
		Price: 100,
		SKU:   "abs-acs-def",
	}
	err := p.validate()
	if err != nil {
		println(err)
		t.Fatal(err)
	}
}
