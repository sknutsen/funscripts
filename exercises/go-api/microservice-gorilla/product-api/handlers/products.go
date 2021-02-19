package handlers

import (
	"context"
	"fmt"
	"net/http"

	"../data"
	"github.com/hashicorp/go-hclog"
)

// Products is a simple handler
type Products struct {
	l         hclog.Logger
	productDB *data.ProductsDB
}

// KeyProduct ...
type KeyProduct struct{}

// NewProducts creates a new products handler
func NewProducts(l hclog.Logger, pdb *data.ProductsDB) *Products {
	return &Products{l, pdb}
}

// MiddlewareProductValidation ...
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Debug("[Error] middleware error")
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Debug("[Error] validating product")
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
