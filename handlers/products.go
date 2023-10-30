package handlers

import (
	"go-microservices/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct{
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products{
	return &Products{logger: logger}
}

func (product *Products) ServeHTTP(res http.ResponseWriter, req *http.Request){
	product.logger.Printf("METHOD: %v", req.Method)
	if req.Method == http.MethodGet{
		product.getProducts(res, req)
		return
	} else if req.Method == http.MethodPost{
		product.postProduct(res, req)
	} else if req.Method == http.MethodPut{
		uri := req.URL.Path
		rx := regexp.MustCompile(`/([0-9]+)`)
		group := rx.FindAllStringSubmatch(uri, -1)
		id, err := strconv.Atoi(group[0][1])

		if err != nil{
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}
		product.putProduct(id, res, req)
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (product *Products) getProducts(res http.ResponseWriter, req *http.Request){
	products := data.GetProducts()
	err := products.ToJson(res)
	if err != nil{
		http.Error(res, "Unable to encode error", http.StatusInternalServerError)
		return
	}
}

func (p *Products) postProduct(res http.ResponseWriter, req *http.Request){
	product := &data.Product{}
	err := product.FromJson(req.Body)
	if err != nil{
		http.Error(res, "Unable to decode error", http.StatusBadRequest)
	}

	data.AddProduct(product)
}
func (p *Products) putProduct(id int, res http.ResponseWriter, req *http.Request){
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