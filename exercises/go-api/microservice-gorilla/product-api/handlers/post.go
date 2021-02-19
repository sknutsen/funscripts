package handlers

import (
	"net/http"

	"../data"
)

// swagger:route POST /products products postProduct
// adds a new product to the collection
// responses:
//	200: noContent

// PostProduct ...
func (p *Products) PostProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}
