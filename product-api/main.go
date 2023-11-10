package main

import (
	"context"
	"go-microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main(){
	// HandleFunc takes my func, creates an HTTP handler, and adds it to the DefaultServeMux
	logger := log.New(os.Stdout, "go-microservices_", log.LstdFlags)
	productHandler := handlers.NewProducts(logger)

	sm := mux.NewRouter()
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", productHandler.PostProduct)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.PutProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", productHandler.DeleteProduct)

	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerHandler := middleware.Redoc(options, nil)
	getRouter.Handle("/docs", swaggerHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := &http.Server{
		Addr: ":9091",
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
}