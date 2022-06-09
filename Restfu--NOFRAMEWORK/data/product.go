package data

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Describtion string  `json:"desc"`
	Price       float32 `json:"price" validate:"numeric,gt=0,lt=200,required"`
	SKU         string  `json:"sku" validate:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func AddProduct(p *Product) bool {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
	return true
}

func getNextID() int {
	lp := ProductList[len(ProductList)-1]
	return lp.ID + 1
}

var ProductList = []*Product{
	{
		ID:          1,
		Name:        "woo",
		Describtion: "descy desc ",
		Price:       300.0,
		CreatedOn:   "",
		UpdatedOn:   "",
	},
}

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Product) validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matches := re.FindAllString(fl.Field().String(), -1)
	// if fl.Field().String() == "invalid" {
	// 	return false
	// }
	if len(matches) != 1 {
		return false
	}
	return true
}

func (p *Product) ToJson(rw io.Writer) error {
	e := json.NewEncoder(rw)
	return e.Encode(p)
}

type Products []*Product
