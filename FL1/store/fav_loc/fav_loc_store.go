package fav_loc

import (
	"database/sql"
	"fmt"
	"github.com/Preksha-zs/FL1/models"
	"github.com/Preksha-zs/FL1/store"
	"log"
)
type fav_loc struct{
	db *sql.DB
}
func New(db *sql.DB) store.Fav_loc {
	return &fav_loc{}
}
func (p *fav_loc)InsertFavLoc(favLoc models.Fav_loc) int64 {
	sqlStatement := `INSERT INTO fav_loc (name, lat, long) VALUES ($1, $2, $3) RETURNING fav_loc_id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := p.db.QueryRow(sqlStatement, favLoc.Name, favLoc.Lat, favLoc.Long).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func (p *fav_loc)GetFavLoc(id int64) (models.Fav_loc, error) {
	var favLoc models.Fav_loc

	sqlStatement := `SELECT * FROM fav_loc WHERE fav_loc_id=$1`

	row := p.db.QueryRow(sqlStatement, id)

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

func (p *fav_loc)GetAllFavLoc() ([]models.Fav_loc, error) {
	var favLocs []models.Fav_loc

	// create the select sql query
	sqlStatement := `SELECT * FROM fav_loc`

	// execute the sql statement
	rows, err := p.db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var favLoc models.Fav_loc

		// unmarshal the row object to user
		err = rows.Scan(&favLoc.ID, &favLoc.Name, &favLoc.Lat, &favLoc.Long)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		favLocs = append(favLocs, favLoc)

	}

	return favLocs, err
}

func (p *fav_loc)UpdateFavLoc(id int64, favLoc models.Fav_loc) int64 {

	sqlStatement := `UPDATE fav_loc SET name=$2, lat=$3, long=$4 WHERE fav_loc_id=$1`

	// execute the sql statement
	res, err := p.db.Exec(sqlStatement, id, favLoc.Name, favLoc.Lat, favLoc.Long)

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
func (p *fav_loc)DeleteFavLoc(id int64) int64 {
	sqlStatement := `DELETE FROM fav_loc WHERE fav_loc_id=$1`

	// execute the sql statement
	res, err := p.db.Exec(sqlStatement, id)

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
