package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Responsible for user authentication
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.ConnectDatabase()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	userPresentOnDatabase, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword([]byte(userPresentOnDatabase.Password), user.Password); err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You are now authenticated"))
}
