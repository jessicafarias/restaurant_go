package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("Home")
}

func restaurants(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "index restaurants")
}

func main() {
	fmt.Println("Server is running on port: 8080")
	http.HandleFunc("/", home)

	http.HandleFunc("/restaurants", restaurants)
	log.Fatal(http.ListenAndServe(":8080", nil))
}