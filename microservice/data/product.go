package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Describtion string  `json:"desc"`
	Price       float32 `json:"price"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
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

type Products []*Product
