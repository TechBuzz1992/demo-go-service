package handlers

import (
	"log"
	"net/http"

	"github.com/demo/demo-go-service/data"
)

//Products struct for creating the handler for the ServeHTTP Handler
type Products struct {
	l *log.Logger
}

//NewProducts as a products fucntion returning the handler instance for handleing the http requests
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	//catch all methods
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle get products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to add product details", http.StatusBadRequest)
	}

	p.l.Printf("Product details are: %#v", prod)

}
