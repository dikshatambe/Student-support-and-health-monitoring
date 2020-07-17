package usergroup

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type usergroupRepository struct {
	conn *sql.DB
}

func NewUsergroupRepository(conn *sql.DB) *usergroupRepository {
	return &usergroupRepository{conn: conn}
}

func (usergroup *usergroupRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.Usergroup)
	return driver.GetById(usergroup.conn, obj, id)
}

func (usergroup *usergroupRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.Usergroup)
	usr := obj.(model.Usergroup)
	// result, err := driver.Create(usergroup.conn, usr)
	result, err := driver.Create(usergroup.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (usergroup *usergroupRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.Usergroup)
	usr := obj.(model.Usergroup)
	// err := driver.UpdateById(usergroup.conn, usr)
	err := driver.UpdateById(usergroup.conn, &usr)
	return obj, err
}

func (usergroup *usergroupRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.Usergroup{Id: id}
	return driver.DeleteById(usergroup.conn, obj, id)
}

func (usergroup *usergroupRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.Usergroup{}
	return driver.GetAll(usergroup.conn, obj, 0, 0)
}

func (usergroup *usergroupRepository) GetByColumnName(cntx context.Context, colname string, value interface{}) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.GroupPermission object repository/grouppermission module")
	obj := &model.Usergroup{}
	return driver.GetByColumnName(usergroup.conn, obj, colname,  value)
}