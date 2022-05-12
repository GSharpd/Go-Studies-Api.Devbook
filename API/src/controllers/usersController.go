package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	userID, err := repo.Create(user)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = userID

	responses.JSONResponse(w, http.StatusOK, user)
}

// Gets all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gettting all users"))
}

// Gets a specific user from the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gettting user"))
}

// Updates user information in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

// Deletes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
