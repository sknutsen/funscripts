package handlers

import (
	"net/http"
	"strconv"

	"../data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products putProduct
// Updates the product with the specified id
// responses:
//	200: noContent

// PutProduct ...
func (p *Products) PutProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
