package user

import (
	"fmt"
	"net/http"
	"webApi/internal/apperror"
	"webApi/internal/handlers"
	"webApi/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	log logging.Logger
}

func NewHandler(log *logging.Logger) handlers.Handler {
	return &handler{
		log: *log,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.getList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.createUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.getUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.updateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.partiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.deleteUser))
}

func (h *handler) getList(w http.ResponseWriter, r *http.Request) error {
	//h.log.Infof("%s %s", r.Method, r.URL.Path)
	//w.WriteHeader(200)
	//w.Write([]byte("this is list of users"))
	return apperror.ErrNotFound
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) error {
	//w.WriteHeader(201)
	//w.Write([]byte("this is create user"))
	return fmt.Errorf("this is API error")
}

func (h *handler) getUserByUUID(w http.ResponseWriter, r *http.Request) error {
	//h.log.Infof("%s %s", r.Method, r.URL.Path)
	//w.WriteHeader(200)
	//w.Write([]byte("this is get user by UUID"))
	return apperror.NewAppError(nil, "test", "test", "test_code")
}

func (h *handler) updateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is update user"))
	return nil
}

func (h *handler) partiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))
	return nil
}

func (h *handler) deleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))
	return nil
}
