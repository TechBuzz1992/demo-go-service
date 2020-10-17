package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product struct for storing products data
type Product struct {
	ID          int     `json:"product-id"`
	Name        string  `json:"product-name"`
	Description string  `json:"product-description"`
	Price       float32 `json:"product-price"`
	SKU         string  `json:"product-sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

//GetProducts returns the list of all products strored in the DB
func GetProducts() Products {
	return productList
}

//Products type for using in the fucntion encoded to JSON
type Products []*Product

//ToJSON function
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

//FromJSON fucntion
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "chocolate",
		Description: "chocolate",
		Price:       125,
		SKU:         "choco",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "plant",
		Description: "plant",
		Price:       250,
		SKU:         "shurbs",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
