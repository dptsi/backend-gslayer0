package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gslayer0/go-middle/controllers/authcontroller"
	"github.com/gslayer0/go-middle/controllers/productcontroller"
	"github.com/gslayer0/go-middle/controllers/usercontroller"
	"github.com/gslayer0/go-middle/middlewares"
	"github.com/gslayer0/go-middle/models"
)

func main() {

	models.ConnectDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/", usercontroller.Welcome)
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/product", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
