package DbConnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "10.10.14.139"
	port     = "1433"
	user     = "usrsprb"
	password = "Banco2014."
	dbname   = "Banca_Movil_GT_3_Desa"
)

//Open Database
func Open() (*sql.DB, error) {
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

//Close Database
func Disconnect(db *sql.DB) {
	db.Close()
}

//Confirm available connection
func Connection() string {
	db, err := Open()
	if err != nil {
		fmt.Print("Something goes wrong")
	}
	Disconnect(db)
	return "Connection available!"
}
