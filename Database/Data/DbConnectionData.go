package DbConnectionData

import (
	"fmt"
	"restaurant_go/Database/Connection"
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