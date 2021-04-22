package DbConnectionData

import (
	"database/sql"
	"fmt"
	"restaurant_go/Database/Connection"
	comment "restaurant_go/Model/Comment.go"
	"restaurant_go/Model/Restaurant"

	_ "github.com/lib/pq"
)

//Generate new Id for Restaurant
func GenerateNewRestaurantId() int {
	DbConnection.Connection();
	db, _ := DbConnection.Open()

	var newId int
	sqlStatement := `
	SELECT MAX(id)+1 FROM restaurants`
	row := db.QueryRow(sqlStatement)
	row.Scan(&newId)

	db.Close()
	return newId
}

//Generate new Id for Comment
func GenerateNewCommentId() int {
	DbConnection.Connection();
	db, _ := DbConnection.Open()

	var newId int
	sqlStatement := `
	SELECT MAX(id)+1 FROM comments`
	row := db.QueryRow(sqlStatement)
	row.Scan(&newId)

	db.Close()
	return newId
}

//Create new Restaurant
func NewRestaurant (name string, description string) {
	db, _ := DbConnection.Open()

	sqlStatement := `
		INSERT INTO restaurants (id, name, description) 
		VALUES ($1, $2, $3)
		RETURNING id`
  id := 0
  err := db.QueryRow(sqlStatement, GenerateNewRestaurantId(), name, description).Scan(&id)
  if err != nil {
    panic(err)
  }
	db.Close();
}

//Create new Comment
func NewComment (){
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
func GetRestaurantById(id string) restaurant.Restaurant{

	db, _ := DbConnection.Open()
  sqlStatement := `SELECT * FROM restaurants WHERE id=$1;`

	var restaurant restaurant.Restaurant
	row := db.QueryRow(sqlStatement, 1)
	err := row.Scan(&restaurant.ID, &restaurant.Nombre, &restaurant.Description)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(restaurant)
	default:
		panic(err)
	}
	db.Close();
	return restaurant
}

//Restaurant index
func GetAllRestaurants() restaurant.AllRestaurants{
	db, _ := DbConnection.Open()
	rows, err := db.Query("SELECT * FROM restaurants LIMIT $1", 3)
	var group restaurant.AllRestaurants
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
		group = append(group, restaurant)
  }
  err = rows.Err()
  if err != nil {
    panic(err)
  }

	return group
}

func GetAllCommentsByRestaurantId(id int){
	db, _ := DbConnection.Open()
	rows, err := db.Query("SELECT * FROM comments WHERE restaurant_id=$1 LIMIT $2", id, 10)
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  for rows.Next() {
    var commentario comment.Comment
    err = rows.Scan(&commentario.Id, &commentario.Username, &commentario.Body, &commentario.RestaurantId)
    if err != nil {
      panic(err)
    }
    fmt.Println(commentario)
  }
  err = rows.Err()
  if err != nil {
    panic(err)
  }
}

