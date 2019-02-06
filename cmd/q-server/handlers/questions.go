package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/3d0c/quiz/cmd/q-server/models/questions"
	"github.com/3d0c/quiz/cmd/q-server/models/sessions"
)

type QuestionsHandler struct{}

func Questions() QuestionsHandler {
	return QuestionsHandler{}
}

func (h QuestionsHandler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		key string = ps.ByName("key")
	)

	if !sessions.Instance().Contains(key) {
		http.Error(w, "Seesion key not found", http.StatusUnauthorized)
		return
	}

	newJsonResponder(w).Write(questions.Instance())
}
