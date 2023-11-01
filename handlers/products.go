package handlers

import (
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
	product := &data.Product{}
	err := product.FromJson(req.Body)
	if err != nil{
		http.Error(res, "Unable to decode error", http.StatusBadRequest)
	}

	data.AddProduct(product)
}
func (p *Products) PutProduct(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id, parseError := strconv.Atoi(vars["id"])
	if parseError != nil{
		http.Error(res, "Unable to convert an id", http.StatusBadRequest)
		return
	}
	product := &data.Product{}
	err := product.FromJson(req.Body)
	p.logger.Printf("Product: %#v", product)
	if err != nil{
		http.Error(res, "Unable to decode error", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(product, id)
	if err == data.ErrProductNotFound{
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil{
		http.Error(res, "Internal Status Error", http.StatusInternalServerError)
		return
	}
	
}