package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	r := newRouter()
	srv := &http.Server {
		Handler: 	r, 
		Addr: 		"127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	go somethingManager()
	go somethingElseManager()

	log.Fatal(srv.ListenAndServe())

}

func somethingManager() {
	fmt.Println("Manager Starting Up")
}

func somethingElseManager() {
	fmt.Println("Manager Else Starting Up")
}