package main

import (
	"fmt"
	"database/sql"
	ps "github.com/Preksha-zs/FL1/service/fav_loc"
	pst "github.com/Preksha-zs/FL1/store/fav_loc"
	_ "github.com/lib/pq"
)
func createConnection() *sql.DB {

	// Open the connection
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=test_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
func main(){
	db:=createConnection()
 	store:=pst.New(db)
 	service:=ps.New(store)
 	fmt.Println(service)
 }
