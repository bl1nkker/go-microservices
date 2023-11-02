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

type Products struct{
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products{
	return &Products{logger: logger}
}

func (product *Products) GetProducts(res http.ResponseWriter, req *http.Request){
	products := data.GetProducts()
	err := products.ToJson(res)
	if err != nil{
		http.Error(res, "Unable to encode error", http.StatusInternalServerError)
		return
	}
}

func (p *Products) PostProduct(res http.ResponseWriter, req *http.Request){
	product := req.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}
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