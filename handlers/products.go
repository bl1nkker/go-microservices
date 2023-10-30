package handlers

import (
	"go-microservices/data"
	"log"
	"net/http"
)

type Products struct{
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products{
	return &Products{logger: logger}
}

func (product *Products) ServeHTTP(res http.ResponseWriter, req *http.Request){
	if req.Method == http.MethodGet{
		product.GetProducts(res, req)
		return
	} else if req.Method == http.MethodPost{
		product.PostProduct(res, req)
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (product *Products) GetProducts(res http.ResponseWriter, req *http.Request){
	products := data.GetProducts()
	err := products.ToJson(res)
	if err != nil{
		http.Error(res, "Unable to marshal error", http.StatusInternalServerError)
	}
}

func (product *Products) PostProduct(res http.ResponseWriter, req *http.Request){
	products := data.GetProducts()
	err := products.ToJson(res)
	if err != nil{
		http.Error(res, "Unable to marshal error", http.StatusInternalServerError)
	}
}