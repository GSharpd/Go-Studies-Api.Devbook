package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err = user.Prepare("signup"); err != nil {
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
	nameOrUserName := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.ConnectDatabase()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	users, err := repo.Get(nameOrUserName)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, users)
}

// Gets a specific user from the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
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

	user, err := repo.GetUserByID(userID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if user.Name == "" {
		err = fmt.Errorf("user with ID %d does not exist", userID)
		responses.ErrorResponse(w, http.StatusNotAcceptable, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, user)
}

// Updates user information in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("you cannot update another user"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
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

	if err = repo.UpdateUser(userID, user); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusNoContent, nil)
}

// Deletes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("you cannot delete another user"))
		return
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	if err = repo.DeleteUser(userID); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusNoContent, nil)
}

// Follows another user using its userID
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	if err = repo.Follow(userID, followerID); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusCreated, nil)
}

// Unfollows another user by its userID
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("you cannot unfollow yourself"))
		return
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)

	if err = repo.Unfollow(userID, followerID); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusNoContent, nil)
}
