package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luisruiz3012/go-gorm-restapi/db"
	"github.com/luisruiz3012/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var user models.User

	db.DB.First(&user, userID)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)

	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var user models.User

	db.DB.First(&user, userID)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	db.DB.Delete(&user) // Stills save the record on DB
	// db.DB.Unscoped().Delete(&user) // Deletes the record on DB

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	var user models.User

	userFound := db.DB.First(&user, userId)
	err := userFound.Error

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewDecoder(r.Body).Decode(&user)

	db.DB.Save(&user)

	json.NewEncoder(w).Encode(&user)
}
