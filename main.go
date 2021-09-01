package main

import (
	"golearn/controllers"
	"golearn/database"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDB()

	log.Println("Starting server at port 8000")
	router := mux.NewRouter().StrictSlash(true)
	initializeRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func initializeRoutes(routes *mux.Router) {
	routes.HandleFunc("/get", controllers.GetUsers).Methods("GET")
}

func initDB() {
	config := database.Config{
		DB_NAME:     "vinnya",
		DB_USERNAME: "vinnyadev",
		DB_PASSWORD: "vinnya1234",
	}

	connectionString := database.GetConnectionString(config)

	log.Println(connectionString)
	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}

	log.Println("connection successful")

}
