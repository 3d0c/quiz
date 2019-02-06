package sessions

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	s := Instance()

	uniq := map[string]bool{}

	for i := 0; i < 1000; i++ {
		key, err := s.generateKey()
		if err != nil {
			t.Fatalf("Error generating key - %s\n", err)
		}

		if n := len(key); n != KeyLength*2 {
			t.Fatalf("Expected key length is %d, obtained: %d\n", KeyLength*2, n)
		}

		if _, ok := uniq[key]; ok {
			t.Fatalf("Non uniq key generated - %s\n", key)
		}
	}
}

func TestSessions(t *testing.T) {
	s := Instance()

	key, err := s.Create()
	if err != nil {
		t.Fatalf("Error creating session - %s\n", err)
	}

	scores := []int{0, 1, 2, 3, 4, 5}

	for _, expected := range scores {
		obtained, err := s.GetScore(key)
		if err != nil {
			t.Fatalf("Error getting score - %s\n", err)
		}
		if expected != obtained {
			t.Fatalf("Expected score '%d', obtained score '%d'\n", expected, obtained)
		}

		obtained, err = s.IncrementScore(key)
		if err != nil {
			t.Fatalf("Error incrementing score - %s\n", err)
		}
		if obtained != expected+1 {
			t.Fatalf("Expected score '%d', obtained score '%d'\n", expected+1, obtained)
		}
	}
}

func TestCompareScore(t *testing.T) {
	s := Instance()
	s.data = nil
	s.data = make(map[string]*session)

	key, err := s.Create()
	if err != nil {
		t.Fatalf("Error creating session - %s\n", err)
	}

	s.data["xxx"] = &session{score: 20}
	s.data["yyy"] = &session{score: 30}
	s.data["zzz"] = &session{score: 0}
	s.data["zz1"] = &session{score: 40}

	testCases := []struct {
		score    int
		expected int
	}{
		{20, 25},
		{0, 0},
		{50, 100},
	}

	for _, cs := range testCases {
		s.data[key].score = cs.score

		obtained, err := s.CompareScore(key)
		if err != nil {
			t.Fatalf("%s\n", err)
		}

		if cs.expected != obtained {
			t.Fatalf("Expected - %d, obtained - %d\n", cs.expected, obtained)
		}
	}
}
