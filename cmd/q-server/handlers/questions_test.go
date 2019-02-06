package handlers

import (
	"testing"

	"github.com/3d0c/quiz/cmd/q-server/models/questions"
	"github.com/3d0c/quiz/pkg/rpc"
)

func TestQuestionsGet(t *testing.T) {
	key = sessionCreate(t)
	endpoint := "http://" + listenOn + "/questions" + "/" + key

	type tmp struct {
		Id string
	}

	resp, err := rpc.Request("GET", endpoint, []byte{}, nil)
	if err != nil {
		t.Fatalf("Error requesting %s - %s\n", endpoint, err)
	}

	result := []tmp{}

	if err = decode(resp.Body, &result); err != nil {
		t.Fatalf("Error reading response - %s\n", err)
	}

	if n := len(result); n != len(questions.Instance()) {
		t.Fatalf("Expected lenght of questions - %d, obtained - %d\n", len(questions.Instance()), n)
	}
}
