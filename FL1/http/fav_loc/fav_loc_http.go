package fav_loc

import (
	"fmt"
	"github.com/Preksha-zs/FL1/models"
	"github.com/Preksha-zs/FL1/service"
	"gitlab.kroger.com/platform/krogo/pkg/krogo"
	"strconv"
)

type Fav_loc struct {
	service service.Fav_loc_service
}

func New(s service.Fav_loc_service) *Fav_loc {
	return &Fav_loc{service: s}
}
func (p *Fav_loc) Create(ctx *krogo.Context) (interface{}, error) {
	fmt.Println("yyyyyyyyyyyyyyyyyyqyqyqyqyyqyqyqyqyyqyqyqyyqyqyqyqy")
	var req *models.Fav_loc
	err := ctx.Bind(&req)
	if err != nil || req == nil {
		return nil, nil
	}
	resp := p.service.Create(req)
	fmt.Println("22222222223333333333")
	//	fmt.Println(resp)
	return resp, nil
}
func (p *Fav_loc) Delete(ctx *krogo.Context) (interface{}, error) {
	ID := ctx.PathParam("id")
	s, _ := strconv.Atoi(ID)
	err := p.service.Delete(int64(s))
	return nil, err
}
func (p *Fav_loc) Update(ctx *krogo.Context) (interface{}, error) {
	var (
		resp *models.Fav_loc
		err  error
	)
	ID := ctx.PathParam("id")
	s, _ := strconv.Atoi(ID)
	var fl *models.Fav_loc
	fmt.Println("im in http update function")
	resp, _ = p.service.Update(int64(s), fl)
	if err != nil {
		return nil, err
	}
	fmt.Println("im in http update function")
	return resp, nil
}
func (p *Fav_loc) Read(ctx *krogo.Context) (interface{}, error) {
	//fmt.Println("%v\n", ctx)
	var (
		resp models.Fav_loc
		err  error
	)
	ID := ctx.PathParam("id")
	s, _ := strconv.Atoi(ID)
	resp, err = p.service.GetByID(int64(s))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
