package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	protos "../../../grpc/currency/protos/currency"
	"../data"
	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// returns a list of products
// responses:
//	200: productsResponse

// GetProducts ...
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("Handle GET products")

	lp, err := p.productDB.GetProducts("GBP")
	if err != nil {
		return
	}

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// returns a list of products
// responses:
//	200: productResponse

// GetProduct ...
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}
	p.l.Debug("Handle GET products")

	prod, err := data.GetProduct(id)
	if err != nil {
		http.Error(rw, fmt.Sprintf("No product found with id: %d", id), http.StatusNotFound)
	}

	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["NOK"]),
		Destination: protos.Currencies(protos.Currencies_value["ISK"]),
	}

	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		http.Error(rw, "Error lol", http.StatusInternalServerError)
	}

	prod.Price = prod.Price * resp.Rate
	err = prod.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
