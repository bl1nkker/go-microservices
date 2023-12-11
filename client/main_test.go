package main

import (
	"client/client"
	"client/client/products"
	"testing"
)

func TestHTTPClient(t *testing.T){
	cfg := client.DefaultTransportConfig().WithHost("localhost:9091")
	client := client.NewHTTPClientWithConfig(nil, cfg)
	productParams := products.NewListProductsParams()
	result, err := client.Products.ListProducts(productParams)
	if err != nil{
		t.Fatal(err)
	}
	
}