package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//GoodBye struct for handling the GoodBye requests
type GoodBye struct {
	l *log.Logger
}

//NewGoodBye Handler
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (h *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("GoodBye!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "GoodBye %s", d)
}
