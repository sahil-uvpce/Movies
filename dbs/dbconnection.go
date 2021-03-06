package dbs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //for postgres driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sahiluvpce"
	dbname   = "movie"
)

// database variable
var DB *sql.DB
var err error

//GetDB get the database conection
func init() {
	if DB == nil {
		connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		DB, err = sql.Open("postgres", connectionString)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("You connected to your database.")
	//return DB
}

//CloseDB close the database conection
func CloseDB(db *sql.DB) {
	err = db.Close()
	fmt.Println("You Closed to your database.")
	if err != nil {
		fmt.Println(err.Error())
	}
}
