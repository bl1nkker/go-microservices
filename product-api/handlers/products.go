// Package classification Product API.
//
// the purpose of this application is to provide an application
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	 Produces:
//	 - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"go-microservices/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// A list of products returns in the response
// swagger:response productsResponse
type ProductsResponseWrapper struct {
	// Description of the response
	// in: body
	Body []data.Product
}

// Product returns in the response
// swagger:response productResponse
type ProductResponseWrapper struct {
	// Description of the response
	// in: body
	Body data.Product
}

// swagger:parameters putProduct postProduct
type ProductParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:response noContent
type NoContentWrapped struct {
}

// swagger:parameters putProduct deleteProduct
type ProductIDParameterWrapper struct {
	// The ID of the product in the database
	// in: path
	// required: true
	ID int `json:"id"`
}

type Products struct{
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products{
	return &Products{logger: logger}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type ErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type ErrorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// swagger:route GET /products products listProducts
// Returns a list of Products
// responses:
//	200: productsResponse
func (product *Products) GetProducts(res http.ResponseWriter, req *http.Request){
	// res.Header().Add("Content-Type", "application/json")
	products := data.GetProducts()
	err := products.ToJson(res)
	if err != nil{
		http.Error(res, "Unable to encode error", http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /products products postProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) PostProduct(res http.ResponseWriter, req *http.Request){
	product := req.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}

// swagger:route PUT /products/{id} products putProduct
// Updates Product by ID
// responses:
//	201: productResponse
func (p *Products) PutProduct(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id, parseError := strconv.Atoi(vars["id"])
	if parseError != nil{
		http.Error(res, "Unable to convert an id", http.StatusBadRequest)
		return
	}

	product := req.Context().Value(KeyProduct{}).(data.Product)
	err := data.UpdateProduct(&product, id)
	if err == data.ErrProductNotFound{
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil{
		http.Error(res, "Internal Status Error", http.StatusInternalServerError)
		return
	}
}

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes Product by ID
// responses:
//	204: noContent
func (p *Products) DeleteProduct(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id, parseError := strconv.Atoi(vars["id"])
	if parseError != nil{
		http.Error(res, "Unable to convert an id", http.StatusBadRequest)
		return
	}
	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound{
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil{
		http.Error(res, "Internal Status Error", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		prod := data.Product{}

		err := prod.FromJson(req.Body)
		if err != nil{
			http.Error(res, "Unable to decode error", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil{
			p.logger.Printf("Validation Error: %s", err)
			http.Error(res, fmt.Sprintf("Validation Error: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), KeyProduct{}, prod)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}