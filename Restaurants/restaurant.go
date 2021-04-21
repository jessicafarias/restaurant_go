package restaurant

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