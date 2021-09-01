package opinions


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
}
Ctrl + Shift + I
*/

//Define Restaurant
type Opinion struct {
	Id           int    `json:"Id"`
	RestaurantId int    `json:"RestaurantId"`
	Body         string `json:"Body"`
	Username     string `json:"Username"`
}

//Define a collection of Restaurants
type Opinions []Opinion
