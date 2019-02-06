package handlers

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/3d0c/quiz/cmd/q-server/models/sessions"
	"github.com/3d0c/quiz/pkg/rpc"
)

const listenOn = "127.0.0.1:6677"

var key string

func TestMain(m *testing.M) {
	router := SetupRouter()

	go func() {
		log.Fatalln(
			http.ListenAndServe(listenOn, router),
		)
	}()

	os.Exit(m.Run())
}

func sessionCreate(t *testing.T) string {
	endpoint := "http://" + listenOn + "/sessions"

	resp, err := rpc.Request("POST", endpoint, []byte{}, nil)
	if err != nil {
		t.Fatalf("Error requesting %s - %s\n", endpoint, err)
	}

	result := SessionsHandler{}

	if err = decode(resp.Body, &result); err != nil {
		t.Fatalf("Error reading response - %s\n", err)
	}

	if n := len(result.Key); n != sessions.KeyLength*2 {
		t.Fatalf("Expected key length - %d, obtained - %d, key - '%s'", n, sessions.KeyLength*2, result.Key)
	}

	return result.Key
}

func TestSessionsCreate(t *testing.T) {
	key = sessionCreate(t)
}

func TestSessionDelete(t *testing.T) {
	key = sessionCreate(t)
	endpoint := "http://" + listenOn + "/sessions" + "/" + key

	_, err := rpc.Request("DELETE", endpoint, []byte{}, nil)
	if err != nil {
		t.Fatalf("Error requesting %s - %s\n", endpoint, err)
	}

	if sessions.Instance().Contains(key) {
		t.Fatalf("Key '%s' still exists\n", key)
	}
}
