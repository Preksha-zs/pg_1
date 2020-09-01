package fav_loc

import (
	"github.com/Preksha-zs/FL1/models"
	"github.com/Preksha-zs/FL1/service"
	"github.com/Preksha-zs/FL1/store"
)

type Fav_loc_service struct {
	store               store.Fav_loc
}

func New(s store.Fav_loc) service.Fav_loc {
	return &Fav_loc_service{
		store: s,
	}
}

func (p *Fav_loc_service) Create(id int64,) int64 {
	//k := models.Fav_loc{}
	//if err := p.store.InsertFavLoc(models.Fav_loc); err != nil {
	//	return nil, err
	//}
   // var resp=*models.Fav_loc
	 var k = p.store.InsertFavLoc(models.Fav_loc{})
	//if err!= nil {
//		return nil, err
	//}
	//}

	return k
}

func (p *Fav_loc_service) GetByID( id int64) (models.Fav_loc, error) {

	var resp, err = p.store.GetFavLoc(int64(id))
	if err != nil {
		return models.Fav_loc{}, err
	}
	return resp, nil
}

func (p *Fav_loc_service)Get() ([]models.Fav_loc, error) {
	//var f []models.Fav_loc
	//var k =[]models.Fav_loc{}
   var f,err=p.store.GetAllFavLoc()
	if err != nil {
			return nil, err
		}
	return f,nil
}

func (p *Fav_loc_service)Update(id int64) (int64, error) {
	//resp:=models.Fav_loc{}
	 var resp = p.store.UpdateFavLoc(int64(id), models.Fav_loc{})
//	if err != nil {
	//	return nil, err
	//}
	return resp, nil
}
func (p *Fav_loc_service)Delete(id int64) error {
	p.store.DeleteFavLoc(int64(id))
	//if err != nil {
	//	return err
	//}
	return nil
}
