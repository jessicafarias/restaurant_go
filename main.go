package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"restaurant_go/Restaurant"
)

var restaurants = restaurant.AllRestaurants{
	{
		ID:          "1",
		Nombre:       "TacoQueto",
		Description: "Deliciosos taquitos",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createRestaurant(w http.ResponseWriter, r *http.Request) {
	var newRestaurant restaurant.Restaurant
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the restaurant title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newRestaurant)
	restaurants = append(restaurants, newRestaurant)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newRestaurant)
}

func getRestaurantById(w http.ResponseWriter, r *http.Request) {
	restaurantID := mux.Vars(r)["id"]

	for _, singleRestaurant := range restaurants {
		if singleRestaurant.ID == restaurantID {
			json.NewEncoder(w).Encode(singleRestaurant)
		}
	}
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(restaurants)
}

func updateRestaurantById(w http.ResponseWriter, r *http.Request) {
	restaurantID := mux.Vars(r)["id"]
	var updatedRestaurant restaurant.Restaurant

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the restaurant title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedRestaurant)

	for i, singleRestaurant := range restaurants {
		if singleRestaurant.ID == restaurantID {
			singleRestaurant.Nombre = updatedRestaurant.Nombre
			singleRestaurant.Description = updatedRestaurant.Description
			restaurants = append(restaurants[:i], singleRestaurant)
			json.NewEncoder(w).Encode(singleRestaurant)
		}
	}
}

func deleteRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID := mux.Vars(r)["id"]

	for i, singleRestaurant := range restaurants {
		if singleRestaurant.ID == restaurantID {
			restaurants = append(restaurants[:i], restaurants[i+1:]...)
			fmt.Fprintf(w, "The restaurant with ID %v has been deleted successfully", restaurantID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/restaurant", createRestaurant).Methods("POST")
	router.HandleFunc("/restaurants", getRestaurants).Methods("GET")
	router.HandleFunc("/restaurant/{id}", getRestaurantById).Methods("GET")
	router.HandleFunc("/restaurant/{id}", updateRestaurantById).Methods("PATCH")
	router.HandleFunc("/restaurant/{id}", deleteRestaurant).Methods("DELETE")
	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}