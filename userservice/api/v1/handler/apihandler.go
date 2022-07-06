package handler

import (
	"net/http"
	"userservice/userservice/errors"
	"userservice/userservice/respond"

	"github.com/gorilla/context"
)

type Handler func(w http.ResponseWriter, r *http.Request) *errors.AppError

func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	// clear gorilla context
	defer context.Clear(r)
	if err != nil {
		// APP Level Error
		// TODO: handle 5XX, notify developers. Configurable
		respond.Fail(w, err)
	}
}
