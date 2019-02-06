package questions

import (
	"testing"
)

func TestQuestionsModel(t *testing.T) {
	if err := Load("../../questions.json"); err != nil {
		t.Fatalf("Error reading questions source file - %s\n", err)
	}

	q := Instance()
	if n := len(q); n != 4 {
		t.Fatalf("Expected length - 4, obtained - %d\n", n)
	}
}

func TestAnswerIsRight(t *testing.T) {
	if err := Load("../../questions.json"); err != nil {
		t.Fatalf("Error reading questions source file - %s\n", err)
	}

	q, ok := Instance().GetQuestion(1)
	if !ok {
		t.Fatalf("Expected question with id 1 not found\n")
	}

	a, ok := q.GetAnswer(1)
	if !ok {
		t.Fatalf("Expected answer with id 1 not found\n")
	}

	if !a.IsRight() {
		t.Fatalf("Expected answer #1 is right, obtained - false\n")
	}
}
