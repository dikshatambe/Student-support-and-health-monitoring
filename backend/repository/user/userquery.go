package user

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type userqueryRepository struct {
	conn *sql.DB
}

func NewUserQueryRepository(conn *sql.DB) *userqueryRepository {
	return &userqueryRepository{conn: conn}
}

func (userquery *userqueryRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.UserQuery)
	return driver.GetById(userquery.conn, obj, id)
}

func (userquery *userqueryRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.UserQuery)
	usr := obj.(model.UserQuery)
	// result, err := driver.Create(userquery.conn, usr)
	result, err := driver.Create(userquery.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (userquery *userqueryRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.UserQuery)
	usr := obj.(model.UserQuery)
	// err := driver.UpdateById(userquery.conn, usr)
	err := driver.UpdateById(userquery.conn, &usr)
	return obj, err
}

func (userquery *userqueryRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.UserQuery{Id: id}
	return driver.DeleteById(userquery.conn, obj, id)
}

func (userquery *userqueryRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.UserQuery{}
	return driver.GetAll(userquery.conn, obj, 0, 0)
}

func (userquery *userqueryRepository) Login(cntx context.Context, id int64, password string) (interface{}, error) {
	obj := new(model.UserQuery)
	return driver.Login(userquery.conn, obj, id, password)
}

func (userquery *userqueryRepository) GetByColumnName(cntx context.Context, colname string, value interface{}) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.GroupPermission object repository/grouppermission module")
	obj := &model.UserQuery{}
	return driver.GetByColumnName(userquery.conn, obj, colname,  value)
}