package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/pucsd2020-pp/ACL/backend/handler"
	"github.com/pucsd2020-pp/ACL/backend/model"
	"github.com/pucsd2020-pp/ACL/backend/repository"
	"github.com/pucsd2020-pp/ACL/backend/repository/files"
)

type Files struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewFilesHandler(conn *sql.DB) *Files {
	return &Files{
		repo: files.NewFilesRepository(conn),
	}
}

func (files *Files) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "files/{id}", Func: files.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "files", Func: files.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "files/{id}", Func: files.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "files/{id}", Func: files.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "files", Func: files.GetAll},
	}
}

func (files *Files) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = files.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (files *Files) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.File
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = files.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (files *Files) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.File{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in files id for tracking update
		// usr.UpdatedBy = 0

		iUsr, err = files.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.File)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (files *Files) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = files.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "Files deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (files *Files) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := files.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
