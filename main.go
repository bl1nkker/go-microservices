package main

import (
	"go-microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main(){
	// HandleFunc takes my func, creates an HTTP handler, and adds it to the DefaultServeMux
	logger := log.New(os.Stdout, "go-microservices_", log.LstdFlags)

	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	// http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Goodbye World")
	// })

	// DefaultServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of 
	// registered patterns and calls the handler for the pattern that most closely matches the URL.
	http.ListenAndServe(":9090", sm)
}