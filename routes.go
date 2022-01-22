package main

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", BaseHandler)
	r.HandleFunc("/test", TestHandler)
	return r
}
