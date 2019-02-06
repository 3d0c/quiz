package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/3d0c/quiz/cmd/q-server/handlers"
	"github.com/3d0c/quiz/cmd/q-server/models/questions"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var (
		listenOn string
		qsource  string
	)

	flag.StringVar(&listenOn, "listen-on", ":5560", "listen on")
	flag.StringVar(&qsource, "q", "questions.json", "question list")
	flag.Parse()

	if err := questions.Load(qsource); err != nil {
		log.Fatalln(err)
	}

	router := handlers.SetupRouter()

	log.Printf("QUIZ server is listening on %s\n", listenOn)

	log.Fatalln(
		http.ListenAndServe(listenOn, router),
	)
}
