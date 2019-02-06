### QUIZ client/server

- build and run `cmd/q-server`
- build and optionally install `cmd/q-client`
- to see available client commands just run

  ```sh
  ./q-client
  
  # expected output
  Usage: ./q-client <COMMAND>
    Available commands:
      quiz.answers
      quiz.end
      quiz.match
      quiz.questions
      quiz.score
      quiz.start
  ```
  
  Each command has it's own description, run it with `--help` option
  
  ```sh
  ./q-client quiz.answers --help
  
  # expected output
  Usage: ./q-client quiz.answers [OPTIONS]

    Description:
    	Answer QUIZ question(s).

    Examples:
    	; answer question id=1 with answer id=1
    	q-client quiz.answers -q=1:1

    	; answer multiple questions
    	q-client quiz.answers -q=1:1 -q=2:1 -q=3:1

    Options:
      -addr=127.0.0.1:5560  quiz server location
      -n=true               dry run, see specific command description
      -q=                   qa in format qid:aid
      -silent=false         suppress output
  ```
  
- please find API endpoints and some description in [router.go](cmd/q-server/handlers/router.go)
- please note, that instead of "cobra" there is another simple cli framework used, you can find it [here](https://github.com/3d0c/cli)
