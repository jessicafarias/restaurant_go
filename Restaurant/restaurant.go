package restaurant

//Describer interface describe the struct being passed in
type Describer interface {
	describes() string
}

//This function doesn't care what type ṕass in as long as type "satisfies the interface"
func WriteAbout(d Describer) string {
	return d.describes()
}

// Colection of all restaurants
/*
type AllRestaurants struct {
	Restauraunts []Restaurant
}*/

//Define Restaurant
type Restaurant struct {
	ID          string `json:"ID"`
	Nombre       string `json:"Nombre"`
	Description string `json:"Description"`
}

//Define a collection of Restaurants
type AllRestaurants []Restaurant
