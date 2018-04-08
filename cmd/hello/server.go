package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getamis/sirius/log"
)

func hello(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(rw, "Hello, current time: %s", time.Now())
}

func bye(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprint(rw, "Bye, see you next time")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)


	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error("failed to start server", "err", err)
	}
}