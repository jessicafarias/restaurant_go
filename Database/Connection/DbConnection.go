package DbConnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
const (
	host ="otto.db.elephantsql.com"
	port ="5432"
	user ="wscjqhpz"
	password ="F_nvM0-Ndv1c3iDvnQiRsvSKZWUvGsAJ"
	dbname = "wscjqhpz"
)

func Open () (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
    panic(err)

  }
	err = db.Ping()
	if err != nil {
		panic(err)
  }
	fmt.Println("Successfully connected")
	return db, err

}

func Disconnect(db *sql.DB){
	db.Close()
	fmt.Println("Disconnected")
}


func Connection() string{
  db, err := Open()
	if (err != nil){
		fmt.Print("Something goes wrong")
	}
	Disconnect(db)
	return "Connection available!"
}
