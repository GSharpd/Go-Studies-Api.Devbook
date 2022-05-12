package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewUserRepository(db)
	userID, err := repo.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created | ID: %d", userID)))
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
