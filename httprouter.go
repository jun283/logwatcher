package main

import (
	"github.com/gorilla/mux"
)

func httpsvr() *mux.Router {
	//s := http.NewServeMux()
	//s.HandleFunc("/foo", foo8001)

	//Create Router
	r := mux.NewRouter()

	//use logging Middleware
	r.Use(loggingMiddleware)

	///use auth Middleware
	amw := authenticationMiddleware{}
	amw.Populate()
	if conf.Authen {
		r.Use(amw.Middleware)
	}

	//handler
	r.HandleFunc("/", PingHandler).Methods("GET")
	r.HandleFunc("/log", LogHandler).Methods("GET")

	return r
}
