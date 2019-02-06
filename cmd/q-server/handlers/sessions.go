package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/3d0c/quiz/cmd/q-server/models/sessions"
)

type SessionsHandler struct {
	Key string
}

func Sessions() SessionsHandler {
	return SessionsHandler{}
}

func (h SessionsHandler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		err error
	)

	if h.Key, err = sessions.Instance().Create(); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	newJsonResponder(w).Write(h)
}

func (h SessionsHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err error
		key string = ps.ByName("key")
	)

	if err = sessions.Instance().Delete(key); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
