package DbConnectionData

import (
	"fmt"
	"restaurant_go/Database/Connection"
	"restaurant_go/Model/Restaurant"
	"database/sql"

	_ "github.com/lib/pq"
)


func NewRestaurant (name string){
	DbConnection.Connection();
	db, _ := DbConnection.Open()

	sqlStatement := `
		INSERT INTO restaurants (id, name, description) 
		VALUES ($1, $2, $3)
		RETURNING id`
  id := 0
  err := db.QueryRow(sqlStatement, 4, name, "Taquitazo feliz").Scan(&id)
  if err != nil {
    panic(err)
  }
  fmt.Println("New record ID is:", id)
	db.Close();
}

func NewComment (){
	DbConnection.Connection();
	db, _ := DbConnection.Open()

	sqlStatement := `
		INSERT INTO comments (id, name, description) 
		VALUES ($1, $2, $3)
		RETURNING id`
  id := 0
  err := db.QueryRow(sqlStatement, 3, "Taquiza", "Taquitazo feliz").Scan(&id)
  if err != nil {
    panic(err)
  }
  fmt.Println("New record ID is:", id)

	db.Close();
}

//Restaurant show
func GetRestaurantById(id int) {
	db, _ := DbConnection.Open()
  sqlStatement := `SELECT * FROM restaurants WHERE id=$1;`
	var restaurant restaurant.Restaurant
	row := db.QueryRow(sqlStatement, 1)
	err := row.Scan(&restaurant.ID, &restaurant.Nombre, &restaurant.Description)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(restaurant)
	default:
		panic(err)
	}

	db.Close();
}

//Restaurant index
func GetAllRestaurants() {
	db, _ := DbConnection.Open()
	rows, err := db.Query("SELECT * FROM restaurants LIMIT $1", 3)
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    var restaurant restaurant.Restaurant
    err = rows.Scan(&restaurant.ID, &restaurant.Nombre, &restaurant.Description)
    if err != nil {
      panic(err)
    }
    fmt.Println(restaurant)
  }
  err = rows.Err()
  if err != nil {
    panic(err)
  }
}

func GetAllCommentsByRestaurantId(id int){
	db, _ := DbConnection.Open()
	rows, err := db.Query("SELECT * FROM comments WHERE restaurant_id=$1 LIMIT $2", id, 3)
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    var restaurant restaurant.Restaurant
    err = rows.Scan(&restaurant.ID, &restaurant.Nombre, &restaurant.Description)
    if err != nil {
      panic(err)
    }
    fmt.Println(restaurant)
  }
  err = rows.Err()
  if err != nil {
    panic(err)
  }
}

