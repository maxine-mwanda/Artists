package main

import (
	"api/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":5000"
	router := mux.NewRouter()
	fmt.Println("Hey there")

	router.HandleFunc("/welcome", controllers.Welcome).Methods("GET")
	router.HandleFunc("/artists", controllers.Singers).Methods("GET", "POST")
	router.HandleFunc("/artist/{singer}", controllers.Singer).Methods("GET")

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("unable to start api because ", err)
	}
}
