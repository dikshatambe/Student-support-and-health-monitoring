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

type UserQuery struct {
	handler.HTTPHandler
	repo  repository.IRepository
	//repo1 repository.JRepository
}

func NewUserQueryHandler(conn *sql.DB) *UserQuery {
	return &UserQuery{
		repo:  user.NewUserQueryRepository(conn),
		//repo1: user.NewUserQueryRepository(conn),
	}
}

func (userquery *UserQuery) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userquery/{id}", Func: userquery.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "userquery", Func: userquery.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "userquery/{id}", Func: userquery.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "userquery/{id}", Func: userquery.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "userquery", Func: userquery.GetAll},
		&handler.HTTPHandler{Authenticated: false, Method: http.MethodGet, Path: "allqueriesofuser/{user_id}", Func: userquery.GetByColumnName},
		//&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "login", Func: userquery.Login},
	}
}

func (userquery *UserQuery) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = userquery.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

/*func (userquery *UserQuery) Login(w http.ResponseWriter, r *http.Request) {
	var usr1 model.UserQuery
	err := json.NewDecoder(r.Body).Decode(&usr1)
	id := usr1.Id
	// fmt.Printf("id is %d\n", id)
	password := usr1.Password
	var usr interface{}
	for {
		if nil != err {
			break
		}

		// usr, err = userquery.repo.Login(r.Context(), id, password)
		usr, err = userquery.repo1.Login(r.Context(), id, password)
		break
	}
	if nil != err {
		handler.WriteJSONResponse(w, r, usr, http.StatusNotFound, err)
	} else {
		handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
	}

}*/

func (userquery *UserQuery) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.UserQuery
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = userquery.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userquery *UserQuery) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.UserQuery{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in userquery id for tracking update
		// usr.UpdatedBy = 0

		iUsr, err = userquery.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.UserQuery)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (userquery *UserQuery) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = userquery.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "UserQuery deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (userquery *UserQuery) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := userquery.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (userquery *UserQuery) GetByColumnName(w http.ResponseWriter, r *http.Request) {
	var grp interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	for {
		if nil != err {
			break
		}

		
		if nil != err {
			break
		}
		grp, err = userquery.repo.GetByColumnName(r.Context(),"user_id",id)
		break
	}

	handler.WriteJSONResponse(w, r, grp, http.StatusOK, err)
}