package fav_loc

import (
	"github.com/Preksha-zs/FL1/http"
	"github.com/Preksha-zs/FL1/models"
	"github.com/Preksha-zs/FL1/service"
	"gitlab.kroger.com/platform/krogo/pkg/krogo"
)

type fav_loc struct {
	service service.Fav_loc
}

func New(s service.Fav_loc) http.Fav_loc {
	return &fav_loc{service: s}
}
func (p *fav_loc) Create(ctx *krogo.Context) (interface{}, error) {
	var req *models.Fav_loc
	err := ctx.Bind(&req)
	if err != nil || req == nil {
		return nil, http.InvalidRequestBody()
	}
	resp := p.service.Create(int64(id))
	return resp, nil
}

func (p *fav_loc) Delete(ctx *krogo.Context) (interface{}, error) {
	ID := ctx.PathParam("id")
	err := p.service.Delete(int64(ID))

	return nil, err
}
func (p *fav_loc) Read(ctx *krogo.Context) (interface{}, error) {
	var (
		resp models.Fav_loc
		err  error
	)
	ID := ctx.PathParam("id")
	resp, err = p.service.GetByID(int64(ID))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
