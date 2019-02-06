package handlers

import (
	"encoding/json"
	"testing"

	"github.com/3d0c/quiz/cmd/q-server/models/questions"
	"github.com/3d0c/quiz/cmd/q-server/models/sessions"
	"github.com/3d0c/quiz/pkg/rpc"
)

func TestAnswer(t *testing.T) {
	if err := questions.Load("../questions.json"); err != nil {
		t.Fatalf("Error reading questions source file - %s\n", err)
	}

	key = sessionCreate(t)
	endpoint := "http://" + listenOn + "/answers" + "/" + key

	type answer struct {
		Qid int
		Aid int
	}

	testCases := []struct {
		Answers []answer
		score   int
	}{
		{
			[]answer{{Qid: 1, Aid: 4}}, 0,
		},
		{
			[]answer{{Qid: 2, Aid: 2}}, 1,
		},
		{
			[]answer{{Qid: 3, Aid: 2}, {Qid: 4, Aid: 2}}, 3,
		},
	}

	for _, cs := range testCases {
		payload, err := json.Marshal(cs)
		if err != nil {
			t.Fatalf("%s\n", err)
		}

		_, err = rpc.Post(endpoint, payload, nil)
		if err != nil {
			t.Fatalf("Error requesting %s - %s, %d\n", endpoint, err, err.(*rpc.Error).Code)
		}

		obtained, err := sessions.Instance().GetScore(key)
		if err != nil {
			t.Fatalf("%s\n", err)
		}

		if cs.score != obtained {
			t.Fatalf("Expected score - %d, obtained - %d\n", cs.score, obtained)
		}
	}

}
