package group

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type groupRepository struct {
	conn *sql.DB
}

func NewGroupRepository(conn *sql.DB) *groupRepository {
	return &groupRepository{conn: conn}
}

func (group *groupRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.Group)
	return driver.GetById(group.conn, obj, id)
}

func (group *groupRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.Group)
	usr := obj.(model.Group)
	// result, err := driver.Create(group.conn, usr)
	result, err := driver.Create(group.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (group *groupRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.Group)
	usr := obj.(model.Group)
	// err := driver.UpdateById(group.conn, usr)
	err := driver.UpdateById(group.conn, &usr)
	return obj, err
}

func (group *groupRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.Group{Id: id}
	return driver.DeleteById(group.conn, obj, id)
}

func (group *groupRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.Group{}
	return driver.GetAll(group.conn, obj, 0, 0)
}

func (group *groupRepository) GetByColumnName(cntx context.Context, colname string, value interface{}) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.GroupPermission object repository/grouppermission module")
	obj := &model.Group{}
	return driver.GetByColumnName(group.conn, obj, colname,  value)
}