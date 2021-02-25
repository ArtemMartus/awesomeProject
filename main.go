package main

import (
	"awesomeProject/src"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", src.Handlers()))
}
