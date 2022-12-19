package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hellohandler
type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// implementing that method on Hello instance(which are interface)
func (h *Hello) ServerHttp(rw http.ResponseWriter, r *http.Request) {
	// NOTE:
	// avoid adding/using concret objects inside the handler so that we enough space to add the
	// objects according to the future cases(testability, future database use etc)

	h.l.Println("Hello, World!")
	// reading from the body and anything that implements the io.ReadCloser there we need to use
	// go library
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	// Using log we are getting data in the logs for the developer but not for the user
	// log.Printf("Data %s\n", data)

	// data for the user
	fmt.Fprintf(rw, "Hello %s\n", data)

}
