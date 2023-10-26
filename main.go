package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	// HandleFunc takes my func, creates an HTTP handler, and adds it to the DefaultServeMux

	http.HandleFunc("/", func(res http.ResponseWriter, req*http.Request) {
		// res is an interface that implements the Writer interface
		// req is the actual pointer(!) to the request object
		log.Println("Hello World")
		data, err := io.ReadAll(req.Body)

		if err != nil{
			http.Error(res, "Oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data: %s\n", data)

		fmt.Fprintf(res, "Hello %s", data)
	})
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// DefaultServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of 
	// registered patterns and calls the handler for the pattern that most closely matches the URL.
	http.ListenAndServe(":9090", nil)
}