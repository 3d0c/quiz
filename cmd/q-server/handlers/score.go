package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/3d0c/quiz/cmd/q-server/models/sessions"
)

type ScoreHandler struct {
	Score   int    `json:",omitempty"`
	Message string `json:",omitempty"`
}

func Score() ScoreHandler {
	return ScoreHandler{}
}

func (h ScoreHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err error
		key string = ps.ByName("key")
	)

	if h.Score, err = sessions.Instance().GetScore(key); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	newJsonResponder(w).Write(h)
}

func (h ScoreHandler) Compare(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		err    error
		key    string = ps.ByName("key")
		result int
	)

	if result, err = sessions.Instance().CompareScore(key); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	h.Message = fmt.Sprintf("You were better than %d%% other quizers", result)

	newJsonResponder(w).Write(h)
}
