package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"Restaurants/restaurant"
	
)

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("Home")
}

// Restaurant type
type Restaurant struct {
	Id          int
	Name        string
	Description string
}

// Colection of all restaurants
type AllRestaurants struct {
	Restauraunts []Restaurant
}

func restaurants(writer http.ResponseWriter, request *http.Request) {
	restaurant2 := Restaurant{1, "Restaruant name", "We love tacos"}
	fmt.Fprint(writer, restaurant2)
	writer.WriteHeader(http.StatusOK)
}


func main() {
	fmt.Println("Server is running on port: 8080")
	http.HandleFunc("/", home)

	http.HandleFunc("/restaurants", restaurants)

	restaurant1 := Restaurant{1, "Restaruant name", "We love tacos"}
	bytesres1, _ := json.Marshal(restaurant1)


	var restaurants AllRestaurants

	if err := json.Unmarshal(bytesres1, &restaurants); err != nil {
		fmt.Println("Error parsing json", err)
	}
	os.Stdout.Write(bytesres1)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
