package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Create struct {
	l *log.Logger
}

func NewCreate(l *log.Logger) *Create {
	return &Create{l}
}

func (c *Create) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
