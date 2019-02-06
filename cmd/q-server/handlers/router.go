package handlers

import (
	"github.com/julienschmidt/httprouter"
)

func SetupRouter() *httprouter.Router {
	router := httprouter.New()

	// Start quiz. Returns uniq key.
	router.POST("/sessions", Sessions().Create)

	// End quiz. E.g. DELETE http://quiz-host/sessions/1c755830f3a86402
	router.DELETE("/sessions/:key", Sessions().Delete)

	// Get current score. E.g. GET http://quiz-host/score/1c755830f3a86402
	router.GET("/score/:key", Score().Get)

	// Compare to others. E.g. GET http://quiz-host/score/1c755830f3a86402/match
	router.GET("/score/:key/match", Score().Compare)

	// Retrive questions with answers. E.g. GET http://quiz-host/questions/1c755830f3a86402
	router.GET("/questions/:key", Questions().Get)

	// Answer question. E.g. POST http://quiz-host/questions/1c755830f3a86402
	// Request body supposed to be valid JSON, an array of objects
	// {"Answers": [{"qid": 1, "aid": 2}]}
	// To answer multiple questions at once POST http://quiz-host/questions/1c755830f3a86402
	// {Answers: [{"qid": 1, "aid": 2}, {"qid": 2, "aid": 1}]}
	router.POST("/answers/:key", Answers().Post)

	return router
}
