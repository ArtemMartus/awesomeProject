package main

import (
	"awesomeProject/src"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	router := httprouter.New()
	router.GET("/", src.IndexHandler)
	router.GET("/api/block/:blockId/total", src.BlockTotalHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
