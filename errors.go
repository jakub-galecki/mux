package mux

import (
	"fmt"
	"log"
	"net/http"
)

type Error struct {
	message string
	status  int
}

func (err *Error) defaultErrorHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(err.status)
	_, errno := fmt.Fprint(w, err.message)
	if errno != nil {
		log.Println(errno.Error())
	}
}
