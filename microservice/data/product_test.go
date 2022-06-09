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
