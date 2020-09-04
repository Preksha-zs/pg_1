package fav_loc

import (
	"fmt"
	"github.com/Preksha-zs/FL1/models"
	"github.com/Preksha-zs/FL1/store"
)

type Fav_loc_service struct {
	store store.Fav_loc
}

func New(s store.Fav_loc) *Fav_loc_service {
	return &Fav_loc_service{
		store: s,
	}
}
func (p *Fav_loc_service) Create(fl *models.Fav_loc) *models.Fav_loc {
	fmt.Println("aaaaaaaaaaaaaaahhhhhhhhhhhhhhhhh")

	fmt.Println(fl)

	resp := p.store.InsertFavLoc(fl)
	return resp
}

func (p *Fav_loc_service) GetByID(id int64) (models.Fav_loc, error) {

	var resp, err = p.store.GetFavLoc(int64(id))
	if err != nil {
		return models.Fav_loc{}, err
	}
	return resp, nil
}

func (p *Fav_loc_service) Get() ([]models.Fav_loc, error) {
	var f, err = p.store.GetAllFavLoc()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (p *Fav_loc_service) Update(id int64, fl *models.Fav_loc) (*models.Fav_loc, error) {
	fmt.Println("im in service update function")
	//	var fl *models.Fav_loc
	var resp = p.store.UpdateFavLoc(int64(id), fl)
	fmt.Println("im in hvkj update function")
	return resp, nil
}
func (p *Fav_loc_service) Delete(id int64) error {
	p.store.DeleteFavLoc(int64(id))
	return nil
}
