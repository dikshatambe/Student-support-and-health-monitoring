package files

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/ACL/backend/driver"
	"github.com/pucsd2020-pp/ACL/backend/model"
)

type filesRepository struct {
	conn *sql.DB
}

func NewFilesRepository(conn *sql.DB) *filesRepository {
	return &filesRepository{conn: conn}
}

func (files *filesRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.File)
	return driver.GetById(files.conn, obj, id)
}

func (files *filesRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.Files)
	usr := obj.(model.File)
	// result, err := driver.Create(files.conn, usr)
	result, err := driver.Create(files.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (files *filesRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.Files)
	usr := obj.(model.File)
	// err := driver.UpdateById(files.conn, usr)
	err := driver.UpdateById(files.conn, &usr)
	return obj, err
}

func (files *filesRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.File{Id: id}
	return driver.DeleteById(files.conn, obj, id)
}

func (files *filesRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.File{}
	return driver.GetAll(files.conn, obj, 0, 0)
}

func (files *filesRepository) GetByColumnName(cntx context.Context, colname string, value interface{}) ([]interface{}, error) {
	//log.Printf("Getting context and creating model.GroupPermission object repository/grouppermission module")
	obj := &model.File{}
	return driver.GetByColumnName(files.conn, obj, colname,  value)
}