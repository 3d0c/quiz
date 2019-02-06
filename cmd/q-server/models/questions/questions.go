package questions

import (
	"encoding/json"
	"os"
)

var (
	instance questions
)

type Answer struct {
	Id     int
	Answer string
	right  bool
}

type Question struct {
	Id       int
	Question string
	Answers  []Answer
}

type questions []Question

func Instance() questions {
	return instance
}

func Load(src string) error {
	r, err := os.OpenFile(src, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer r.Close()

	return json.NewDecoder(r).Decode(&instance)
}

func (a *Answer) UnmarshalJSON(b []byte) error {
	type tmp struct {
		Id     int
		Answer string
		Right  bool
	}

	t := tmp{}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	a.Id = t.Id
	a.Answer = t.Answer
	a.right = t.Right

	return nil
}

func (qs questions) GetQuestion(id int) (Question, bool) {
	for i, _ := range qs {
		if qs[i].Id == id {
			return qs[i], true
		}
	}

	return Question{}, false
}

func (q Question) GetAnswer(id int) (Answer, bool) {
	for i, _ := range q.Answers {
		if q.Answers[i].Id == id {
			return q.Answers[i], true
		}
	}

	return Answer{}, false
}

func (a Answer) IsRight() bool {
	return a.right
}
