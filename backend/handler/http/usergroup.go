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
	"github.com/pucsd2020-pp/ACL/backend/repository/usergroup"
)

type Usergroup struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewUsergroupHandler(conn *sql.DB) *Usergroup {
	return &Usergroup{
		repo: usergroup.NewUsergroupRepository(conn),
	}
}

func (usergroup *Usergroup) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "usergroup/{id}", Func: usergroup.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "usergroup", Func: usergroup.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "usergroup/{id}", Func: usergroup.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "usergroup/{id}", Func: usergroup.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "usergroup", Func: usergroup.GetAll},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodGet, Path: "allgroupsofuser/{user_id}", Func: usergroup.GetByColumnName},
	}
}

func (usergroup *Usergroup) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		//usr, err = usergroup.repo.GetByID(r.Context(), id)
		usr, err = usergroup.repo.GetByColumnName(r.Context(), "id",id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (usergroup *Usergroup) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.Usergroup
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = usergroup.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (usergroup *Usergroup) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.Usergroup{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in usergroup id for tracking update
		// usr.UpdatedBy = 0

		iUsr, err = usergroup.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.Usergroup)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (usergroup *Usergroup) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = usergroup.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "Usergroup deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (usergroup *Usergroup) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := usergroup.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}


func (usergroup *Usergroup) GetByColumnName(w http.ResponseWriter, r *http.Request) {
	var grp interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	for {
		if nil != err {
			break
		}

		
		if nil != err {
			break
		}
		grp, err = usergroup.repo.GetByColumnName(r.Context(),"user_id",id)
		break
	}

	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}