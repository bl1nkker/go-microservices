package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// data.Product represents a product
// swagger:model
type Product struct {
    // ID of the product
    // required: true
    // example: 1
    ID int `json:"id"`
    // Name of the product
    // required: true
    // example: Latte
    Name string `json:"name"`
    // Description of the product
    // example: Cool Latte
    Description string `json:"description"`
    // Price of the product
    // required: true
    // example: 12.99
	// min: 1
    // max: 999
    Price float32 `json:"price"`
    // SKU of the product
    // required: true
    // example: abc-def-xyz
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

func (product *Product) Validate() error{
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(product)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
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