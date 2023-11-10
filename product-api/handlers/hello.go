package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct{
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h*Hello) ServeHTTP(res http.ResponseWriter, req *http.Request){
	// res is an interface that implements the Writer interface
	// req is the actual pointer(!) to the request object
	h.logger.Println("Hello World")
	data, err := io.ReadAll(req.Body)

	if err != nil{
		http.Error(res, "Oops", http.StatusBadRequest)
		return
	}
	h.logger.Printf("Data: %s\n", data)

	fmt.Fprintf(res, "Hello %s", data)
}