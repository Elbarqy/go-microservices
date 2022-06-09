package handlers

import (
	"log"
	"microservice/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
