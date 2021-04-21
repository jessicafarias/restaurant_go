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

func Connection() string{
  psqlInfo := fmt.Sprintf("host=%s port=%v user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

	return "Successfully connected!"
}
