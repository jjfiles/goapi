package main

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", BaseHandler)
	return r
}