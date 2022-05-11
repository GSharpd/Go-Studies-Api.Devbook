package controllers

import "net/http"

// Creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new user"))
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
