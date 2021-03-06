package repository

import (
	"context"
)

type IRepository interface {
	GetByID(context.Context, int64) (interface{}, error)
	Create(context.Context, interface{}) (interface{}, error)
	Update(context.Context, interface{}) (interface{}, error)
	Delete(context.Context, int64) error
	GetAll(context.Context) ([]interface{}, error)
	GetByColumnName(context.Context, string, interface{}) ([]interface{}, error)
}

type JRepository interface {
	Login(context.Context, int64, string) (interface{}, error)
}

type Repository struct {
}

func (repo *Repository) GetByID(cntx context.Context, id int64) (obj interface{}, err error) {
	return
}

func (repo *Repository) Create(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (repo *Repository) Update(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

func (repo *Repository) Delete(cntx context.Context, id int64) (deleted bool, err error) {
	return
}

func (repo *Repository) GetAll(cntx context.Context) (obj []interface{}, err error) {
	return
}

func (repo *Repository) Login(cntx context.Context, id int64, password string) (obj interface{}, err error) {
	return
}

func (repo *Repository) GetByColumnName(cntx context.Context, id string) (obj interface{}, err error){
	return
}