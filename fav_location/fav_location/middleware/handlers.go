package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"fav_location/models" // models package where FavLoc schema is defined
	"log"
	"net/http" // used to access the request and response object of the api
	"strconv"  // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
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

func CreateFavLoc(w http.ResponseWriter, r *http.Request) {

	var favLoc models.FavLoc

	err := json.NewDecoder(r.Body).Decode(&favLoc)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := insertFavLoc(favLoc)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "FavLoc created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func GetFavLoc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	favLoc, err := getFavLoc(int64(id))

	if err != nil {
		log.Fatalf("Unable to get fav loc. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(favLoc)
}

func GetAllFavLoc(w http.ResponseWriter, r *http.Request) {

	favLocs, err := getAllFavLoc()

	if err != nil {
		log.Fatalf("Unable to get all fav loc. %v", err)
	}

	json.NewEncoder(w).Encode(favLocs)
}

func UpdateFavLoc(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var favLoc models.FavLoc

	err = json.NewDecoder(r.Body).Decode(&favLoc)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateFavLoc(int64(id), favLoc)

	// format the message string
	msg := fmt.Sprintf("favLoc updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteFavLoc(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteFavLoc(int64(id))

	// format the message string
	msg := fmt.Sprintf("FavLoc updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
func insertFavLoc(favLoc models.FavLoc) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	sqlStatement := `INSERT INTO fav_loc (name, lat, long) VALUES ($1, $2, $3) RETURNING fav_loc_id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, favLoc.Name, favLoc.Lat, favLoc.Long).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func getFavLoc(id int64) (models.FavLoc, error) {
	// create the postgres db connection
	db := createConnection()

	defer db.Close()

	var favLoc models.FavLoc

	sqlStatement := `SELECT * FROM fav_loc WHERE fav_loc_id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&favLoc.ID, &favLoc.Name, &favLoc.Lat, &favLoc.Long)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return favLoc, nil
	case nil:
		return favLoc, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return favLoc, err
}

func getAllFavLoc() ([]models.FavLoc, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var favLocs []models.FavLoc

	// create the select sql query
	sqlStatement := `SELECT * FROM fav_loc`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var favLoc models.FavLoc

		// unmarshal the row object to user
		err = rows.Scan(&favLoc.ID, &favLoc.Name, &favLoc.Lat, &favLoc.Long)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		favLocs = append(favLocs, favLoc)

	}

	return favLocs, err
}

func updateFavLoc(id int64, favLoc models.FavLoc) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE fav_loc SET name=$2, lat=$3, long=$4 WHERE fav_loc_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, favLoc.Name, favLoc.Lat, favLoc.Long)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func deleteFavLoc(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM fav_loc WHERE fav_loc_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
