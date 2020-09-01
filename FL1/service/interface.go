package service

import "github.com/Preksha-zs/FL1/models"

type Fav_loc interface{
	Create(id int64,) int64
	GetByID( id int64) (models.Fav_loc, error)
	Get() ([]models.Fav_loc, error)
	Update(id int64) (int64, error)
	Delete(id int64) error
	//Create(id int64) (*models.Fav_loc, error)
	//GetByID( id int64) (*models.Fav_loc, error)
	//Get() ([]*models.Fav_loc, error)
//	Update(id int64) (*models.Fav_loc, error)
	//Delete(id int64) error
}
