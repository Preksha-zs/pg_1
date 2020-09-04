package main

import (
	"database/sql"
	"fmt"
	ph "github.com/Preksha-zs/FL1/http/fav_loc"
	ps "github.com/Preksha-zs/FL1/service/fav_loc"
	pst "github.com/Preksha-zs/FL1/store/fav_loc"
	_ "github.com/lib/pq"
	"gitlab.kroger.com/platform/krogo/pkg/krogo"
	//"os"
)

func createConnection() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=test_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
func main() {
	k := krogo.New()
	k.Server.HTTP.Port = 9091
	db := createConnection()
	store := pst.New(db)
	service := ps.New(store)
	Handler := ph.New(service)
	fmt.Printf("%v\n", Handler)
	k.POST("/FL1", Handler.Create)
	k.PUT("/FL1/{id}", Handler.Update)
	k.GET("/FL1/{id}", Handler.Read)
	k.DELETE("/FL1/{id}", Handler.Delete)
	k.Start()
}
