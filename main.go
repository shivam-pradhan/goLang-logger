package main

import (
	"github.com/gorilla/mux"
	"logger/controller"

	//"log"
	"net/http"
)

func main() {
	startRouter()
}

// startRouter ...
/**
 * router mapping for each request
 */
func startRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", controller.Hello)
	r.HandleFunc("/logLevel/{logLevel}", controller.UpdateCurrentLogLevel).Methods("PUT")
	r.HandleFunc("/logLevel", controller.GetCurrentLogLevel).Methods("GET")

	http.Handle("/", r)

	err := http.ListenAndServe(":8083", nil)
	checkErr(err)
}

// checkErr ...
/**
 *
 *Check for error and panic the system start if any error while booting occured
 */
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
