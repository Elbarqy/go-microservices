package handlers

import (
	"encoding/json"
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
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(data.ProductList)
	if err != nil {
		p.l.Println("Couldn't marshal this", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	rw.Write(resp)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")
	prod := &data.Product{}
	println(r.Body)
	err := prod.FromJson(r.Body)
	if err != nil {
		p.l.Println("okay this is an error")
		p.l.Println(err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	ok := data.AddProduct(prod)
	if ok {
		resp, _ := json.Marshal(prod)
		rw.Write(resp)
	}
}
