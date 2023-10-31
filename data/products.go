package data

import (
	"encoding/json"
	"fmt"
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

func UpdateProduct(p *Product, id int) error{
	prod, pos, err := findProduct(id)
	if err != nil{
		return err
	}
	prod.ID = id
	productsList[pos] = p
	return nil
}

func getNextId() int {
	lp := productsList[len(productsList) - 1]
	return lp.ID + 1
}


var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error){
	for i, p := range(productsList){
		if p.ID == id{
			return p, i, nil
		}
	}
	return nil, 0, ErrProductNotFound
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
		ID: 2,
		Name: "Cappuchino",
		Description: "Cool Cappuchino",
		Price: 2.99,
		SKU: "abc228",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	{
		ID: 3,
		Name: "Typical Coffee",
		Description: "Cool Typical Coffee",
		Price: 3.99,
		SKU: "abc337",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}