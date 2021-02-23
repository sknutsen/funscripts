// Package classification of product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers

import "../data"

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// A product returns in the response
// swagger:response productResponse
type productResponseWrapper struct {
	// All products in the system
	// in: body
	Body data.Product
}

// swagger:response noContent
type productsNoContent struct {
}

// swagger:parameters deleteProduct listSingleProduct putProduct
type productIDParamsWrapper struct {
	// The id of the product to delete
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters postProduct putProduct
type productParamsWrapper struct {
	// Product data structure
	// in: body
	// required: true
	Body data.Product
}
