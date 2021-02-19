package handlers

import (
	"net/http"
	"strconv"

	"../data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes the product with the specified id
// responses:
//	200: noContent

// DeleteProduct deltes the product with the given id
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
