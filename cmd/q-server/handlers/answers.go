package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/3d0c/quiz/cmd/q-server/models/questions"
	"github.com/3d0c/quiz/cmd/q-server/models/sessions"
)

type AnswersHandler struct {
	Answers []struct {
		Qid int
		Aid int
	}
}

func Answers() AnswersHandler {
	return AnswersHandler{}
}

func (h AnswersHandler) Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		key string = ps.ByName("key")
		err error
	)

	if !sessions.Instance().Contains(key) {
		http.Error(w, "Seesion key not found", http.StatusUnauthorized)
		return
	}

	if err = decode(r.Body, &h); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	for _, answer := range h.Answers {
		answered, err := sessions.Instance().IsAnswered(key, answer.Qid)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if answered {
			http.Error(w, "Question is already answered", http.StatusBadRequest)
			return
		}

		q, ok := questions.Instance().GetQuestion(answer.Qid)
		if !ok {
			http.Error(w, "Question not found", http.StatusNotFound)
			return
		}

		a, ok := q.GetAnswer(answer.Aid)
		if !ok {
			http.Error(w, "Answer not found", http.StatusNotFound)
			return
		}

		if a.IsRight() {
			if _, err = sessions.Instance().IncrementScore(key); err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}

		if err = sessions.Instance().SetAnswered(key, answer.Qid); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}
