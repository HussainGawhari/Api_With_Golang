package main

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", controller.AllEmployee).Methods("GET")
	router.HandleFunc("/employee", controller.AllEmployee).Methods("GET")
	router.HandleFunc("/employee/{id}", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", controller.DeleteEmployee).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
