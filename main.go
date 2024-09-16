package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luisruiz3012/go-gorm-restapi/db"
	"github.com/luisruiz3012/go-gorm-restapi/models"
	"github.com/luisruiz3012/go-gorm-restapi/routes"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	// Users routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", routes.UpdateUser).Methods("PUT")

	// Tasks routes
	r.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTask).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTasks).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", routes.UpdateTask).Methods("PUT")

	http.ListenAndServe(":3000", r)
}
