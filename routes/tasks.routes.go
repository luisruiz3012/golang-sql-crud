package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luisruiz3012/go-gorm-restapi/db"
	"github.com/luisruiz3012/go-gorm-restapi/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	var task models.Task

	db.DB.First(&task, taskId)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdUser := db.DB.Create(&task)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	var task models.Task

	taskFound := db.DB.First(&task, taskId)
	err := taskFound.Error

	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	db.DB.Delete(&task)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	var task models.Task

	foundTaks := db.DB.First(&task, taskId)
	err := foundTaks.Error

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewDecoder(r.Body).Decode(&task)

	db.DB.Save(&task)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}
