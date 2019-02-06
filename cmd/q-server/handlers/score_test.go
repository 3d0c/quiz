package handlers

import (
	"testing"

	"github.com/3d0c/quiz/cmd/q-server/models/sessions"
	"github.com/3d0c/quiz/pkg/rpc"
)

func TestScore(t *testing.T) {
	key = sessionCreate(t)

	testCases := []struct {
		endpoint string
		expected interface{}
	}{
		{"http://" + listenOn + "/score" + "/" + key, 100},
		{"http://" + listenOn + "/score" + "/" + key + "/match", "You were better than 100% other quizers"},
	}

	for i := 1; i <= 100; i++ {
		sessions.Instance().IncrementScore(key)
	}

	for _, cs := range testCases {
		resp, err := rpc.Request("GET", cs.endpoint, []byte{}, nil)
		if err != nil {
			t.Fatalf("Error requesting %s - %s\n", cs.endpoint, err)
		}

		result := ScoreHandler{}

		if err = decode(resp.Body, &result); err != nil {
			t.Fatalf("Error reading response - %s\n", err)
		}

		switch cs.expected.(type) {
		case int:
			if result.Score != cs.expected.(int) {
				t.Fatalf("Expected score - %d, obtained - %d\n", cs.expected.(int), result.Score)
			}
		case string:
			if result.Message != cs.expected.(string) {
				t.Fatalf("Expected message - '%s', obtained - '%s'", cs.expected.(string), result.Message)
			}
		}
	}

}
