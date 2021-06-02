package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"restaurant_go/Database/DbConnectionData"
	comment "restaurant_go/Model/Comment.go"
	"restaurant_go/Model/Restaurant"

	"github.com/gorilla/mux"
)

var restaurants = restaurant.AllRestaurants{}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Restaurant API with Go")
}

// POST restaurant
func createRestaurant(w http.ResponseWriter, r *http.Request) {
	var newRestaurant restaurant.Restaurant

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the restaurant title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newRestaurant)
	DbConnectionData.NewRestaurant(newRestaurant.Nombre, newRestaurant.Description)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newRestaurant)
}

// SHOW /restaurant/1
func getRestaurantById(w http.ResponseWriter, r *http.Request) {
	restaurantID := mux.Vars(r)["id"]

	singleRestaurant := DbConnectionData.GetRestaurantById(restaurantID)
	json.NewEncoder(w).Encode(singleRestaurant)
}

//INDEX /restaurants
func getRestaurants(w http.ResponseWriter, r *http.Request) {
	restaurants = DbConnectionData.GetAllRestaurants()
	json.NewEncoder(w).Encode(restaurants)	
}

// UPDATE restaurant/1 (not implemented)
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

// DELETE restaurant/2 (not implemented)
func deleteRestaurant(w http.ResponseWriter, r *http.Request) {
	restaurantID := mux.Vars(r)["id"]

	for i, singleRestaurant := range restaurants {
		if singleRestaurant.ID == restaurantID {
			restaurants = append(restaurants[:i], restaurants[i+1:]...)
			fmt.Fprintf(w, "The restaurant with ID %v has been deleted successfully", restaurantID)
		}
	}
}

// POST comment
func createComment(w http.ResponseWriter, r *http.Request) {
	// {"Username": "Mildred", "Body": "Body", "RestaurantId":1}

	var newComment comment.Comment
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the comment parameters")
	}

	json.Unmarshal(reqBody, &newComment)
	err = DbConnectionData.NewComment(newComment)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newComment)
}

// SHOW All comments by Restaurant ID
func getComments(w http.ResponseWriter, r *http.Request) {
	restaurantID := mux.Vars(r)["restaurant_id"]
	commentsbyid := DbConnectionData.GetAllCommentsByRestaurantId(restaurantID)
	json.NewEncoder(w).Encode(commentsbyid)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/restaurant", createRestaurant).Methods("POST")
	router.HandleFunc("/restaurants", getRestaurants).Methods("GET")
	router.HandleFunc("/restaurant/{id}", getRestaurantById).Methods("GET")
	router.HandleFunc("/restaurant/{id}", updateRestaurantById).Methods("PATCH")
	router.HandleFunc("/restaurant/{id}", deleteRestaurant).Methods("DELETE")

	router.HandleFunc("/comment", createComment).Methods("POST")
	router.HandleFunc("/comments/{restaurant_id}", getComments).Methods("GET")
	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
