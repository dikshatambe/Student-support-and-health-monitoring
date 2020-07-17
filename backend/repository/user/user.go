package user

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *userRepository {
	return &userRepository{conn: conn}
}

func (user *userRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.User)
	return driver.GetById(user.conn, obj, id)
}

func (user *userRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.User)
	usr := obj.(model.User)
	// result, err := driver.Create(user.conn, usr)
	result, err := driver.Create(user.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user *userRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.User)
	usr := obj.(model.User)
	// err := driver.UpdateById(user.conn, usr)
	err := driver.UpdateById(user.conn, &usr)
	return obj, err
}

func (user *userRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.User{Id: id}
	return driver.DeleteById(user.conn, obj, id)
}

func (user *userRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.User{}
	return driver.GetAll(user.conn, obj, 0, 0)
}

func (user *userRepository) Login(cntx context.Context, id int64, password string) (interface{}, error) {
	obj := new(model.User)
	return driver.Login(user.conn, obj, id, password)
}

func (user *userRepository) GetByColumnName(cntx context.Context, colname string, value interface{}) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.GroupPermission object repository/grouppermission module")
	obj := &model.User{}
	return driver.GetByColumnName(user.conn, obj, colname,  value)
}