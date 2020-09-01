package store

import "github.com/Preksha-zs/FL1/models"

type Fav_loc interface{
	InsertFavLoc(favLoc models.Fav_loc) int64
	GetFavLoc(id int64) (models.Fav_loc, error)
	GetAllFavLoc() ([]models.Fav_loc, error)
	UpdateFavLoc(id int64, favLoc models.Fav_loc) int64
	DeleteFavLoc(id int64) int64
}