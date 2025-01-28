package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	imageHandler "github.com/voyagera/backend/internal/handlers/image"
	loginHandler "github.com/voyagera/backend/internal/handlers/login"
	tripsHandler "github.com/voyagera/backend/internal/handlers/trips"
)

var secretKey = []byte("secret-key")

func main() {
	router := mux.NewRouter()

	// handle login and sign-up
	router.HandleFunc("/v1/auth/", loginHandler.LoginHandler).Methods("POST")

	// handle users

	// handle trips
	router.HandleFunc("/v1/trips/", tripsHandler.TripsHandlerGET).Methods("GET")
	router.HandleFunc("/v1/trips/", tripsHandler.TripsHandlerPOST).Methods("POST")
	router.HandleFunc("/v1/trips/", tripsHandler.TripsHandlerDELETE).Methods("DELETE")
	router.HandleFunc("/v1/trips/", tripsHandler.TripsHandlerPUT).Methods("PUT")

	// handle image
	router.HandleFunc("/v1/image/", imageHandler.ImageHandler).Methods("GET")

	fmt.Println("Starting the server on http://localhost:4000")

	err := http.ListenAndServe("localhost:4000", router)
	if err != nil {
		fmt.Println("Could not start the server", err)
	}
}
