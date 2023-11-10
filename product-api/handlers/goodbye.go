package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct{
	logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye{
	return &Goodbye{logger: logger}
}

func (goodbye *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request){
	goodbye.logger.Println("Goodbye World")
}