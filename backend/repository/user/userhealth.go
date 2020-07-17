package user

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type userhealthRepository struct {
	conn *sql.DB
}

func NewUserHealthRepository(conn *sql.DB) *userhealthRepository {
	return &userhealthRepository{conn: conn}
}

func (userhealth *userhealthRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.UserHealth)
	return driver.GetById(userhealth.conn, obj, id)
}

func (userhealth *userhealthRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.UserHealth)
	usr := obj.(model.UserHealth)
	// result, err := driver.Create(userhealth.conn, usr)
	result, err := driver.Create(userhealth.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (userhealth *userhealthRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.UserHealth)
	usr := obj.(model.UserHealth)
	// err := driver.UpdateById(userhealth.conn, usr)
	err := driver.UpdateById(userhealth.conn, &usr)
	return obj, err
}

func (userhealth *userhealthRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.UserHealth{Id: id}
	return driver.DeleteById(userhealth.conn, obj, id)
}

func (userhealth *userhealthRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.UserHealth{}
	return driver.GetAll(userhealth.conn, obj, 0, 0)
}

func (userhealth *userhealthRepository) Login(cntx context.Context, id int64, password string) (interface{}, error) {
	obj := new(model.UserHealth)
	return driver.Login(userhealth.conn, obj, id, password)
}

func (userhealth *userhealthRepository) GetByColumnName(cntx context.Context, colname string, value interface{}) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.GroupPermission object repository/grouppermission module")
	obj := &model.UserHealth{}
	return driver.GetByColumnName(userhealth.conn, obj, colname,  value)
}