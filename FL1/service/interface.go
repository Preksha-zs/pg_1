package service

import (
	"github.com/Preksha-zs/FL1/models"
)

type Fav_loc_service interface {
	Create(loc *models.Fav_loc) *models.Fav_loc
	GetByID(id int64) (models.Fav_loc, error)
	Get() ([]models.Fav_loc, error)
	Update(id int64, fl *models.Fav_loc) (*models.Fav_loc, error)
	Delete(id int64) error
}
