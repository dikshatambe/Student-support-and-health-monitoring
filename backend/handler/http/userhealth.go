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
	"github.com/pucsd2020-pp/ACL/backend/repository/user"
)

type UserHealth struct {
	handler.HTTPHandler
	repo  repository.IRepository
	repo1 repository.JRepository
}

func NewUserHealthHandler(conn *sql.DB) *UserHealth {
	return &UserHealth{
		repo:  user.NewUserHealthRepository(conn),
		repo1: user.NewUserHealthRepository(conn),
	}
}

func (userhealth *UserHealth) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userhealth/{id}", Func: userhealth.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "userhealth", Func: userhealth.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "userhealth/{id}", Func: userhealth.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "userhealth/{id}", Func: userhealth.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userhealth", Func: userhealth.GetAll},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodGet, Path: "alldetailsofuser/{user_id}", Func: userhealth.GetByColumnName},
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "login", Func: userhealth.Login},
	}
}

func (userhealth *UserHealth) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = userhealth.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

/*func (userhealth *UserHealth) Login(w http.ResponseWriter, r *http.Request) {
	var usr1 model.UserHealth
	err := json.NewDecoder(r.Body).Decode(&usr1)
	id := usr1.Id
	// fmt.Printf("id is %d\n", id)
	password := usr1.Password
	var usr interface{}
	for {
		if nil != err {
			break
		}

		// usr, err = userhealth.repo.Login(r.Context(), id, password)
		usr, err = userhealth.repo1.Login(r.Context(), id, password)
		break
	}
	if nil != err {
		handler.WriteJSONResponse(w, r, usr, http.StatusNotFound, err)
	} else {
		handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
	}

}*/

func (userhealth *UserHealth) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.UserHealth
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = userhealth.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userhealth *UserHealth) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.UserHealth{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in userhealth id for tracking update
		// usr.UpdatedBy = 0

		iUsr, err = userhealth.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.UserHealth)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userhealth *UserHealth) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = userhealth.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "UserHealth deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (userhealth *UserHealth) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := userhealth.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (userhealth *UserHealth) GetByColumnName(w http.ResponseWriter, r *http.Request) {
	var grp interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	for {
		if nil != err {
			break
		}

		
		if nil != err {
			break
		}
		grp, err = userhealth.repo.GetByColumnName(r.Context(),"user_id",id)
		break
	}

	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}