package main

import (
	"context"
	"go-microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	// HandleFunc takes my func, creates an HTTP handler, and adds it to the DefaultServeMux
	logger := log.New(os.Stdout, "go-microservices_", log.LstdFlags)

	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodbye(logger)
	productHandler := handlers.NewProducts(logger)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", productHandler)
	server := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func(){
		err := server.ListenAndServe()
		if err != nil{
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal, 5)
	signal.Notify(signalChannel, os.Interrupt)

	logger.Println("Waiting for a signal...")
	sig := <-signalChannel
	logger.Println("Signal Received!")
	logger.Println("Received terminate. Graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(tc)

	// DefaultServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of 
	// registered patterns and calls the handler for the pattern that most closely matches the URL.
}