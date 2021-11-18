package router

import (
	"log"
	"net/http"
)

func HTTPBadRequest(w http.ResponseWriter, msg string, err error) {
	httpError(w, http.StatusBadRequest, msg, err)
}

func HTTPInternalServerError(w http.ResponseWriter, msg string, err error) {
	httpError(w, http.StatusInternalServerError, msg, err)
}

func httpError(w http.ResponseWriter, httpStatus int, msg string, err error) {
	http.Error(w, msg, httpStatus)
	log.Printf("%T: %#v \n", err, err)
}
