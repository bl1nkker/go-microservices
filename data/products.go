package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Products []*Product

func (products *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}

func (product *Product) FromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

func GetProducts() Products{
	return productsList
}

func AddProduct(product *Product) {
	product.ID = getNextId()
	productsList = append(productsList, product)
}

func getNextId() int {
	lp := productsList[len(productsList) - 1]
	return lp.ID + 1
}

var productsList = []*Product{
	{
		ID: 1,
		Name: "Latte",
		Description: "Cool Latte",
		Price: 12.99,
		SKU: "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID: 1,
		Name: "Cappuchino",
		Description: "Cool Cappuchino",
		Price: 2.99,
		SKU: "abc228",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}