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
		product.getProducts(res, req)
		return
	} else if req.Method == http.MethodPost{
		product.postProduct(res, req)
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (product *Products) getProducts(res http.ResponseWriter, req *http.Request){
	products := data.GetProducts()
	err := products.ToJson(res)
	if err != nil{
		http.Error(res, "Unable to encode error", http.StatusInternalServerError)
	}
}

func (p *Products) postProduct(res http.ResponseWriter, req *http.Request){
	// data := req.Body
	product := &data.Product{}
	p.logger.Printf("Incoming Body: %v", req.Body)
	err := product.FromJson(req.Body)
	p.logger.Printf("Incoming Product: %#v", product)
	if err != nil{
		http.Error(res, "Unable to decode error", http.StatusBadRequest)
	}

	data.AddProduct(product)
}