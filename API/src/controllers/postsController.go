package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post

	if err := json.Unmarshal(requestBody, &post); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	post.PosterID = tokenUserID

	db, err := database.ConnectDatabase()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPostsRepository(db)

	post.ID, err = repo.CreateNewPost(post)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusCreated, post)
}

// Gets user feed posts
func GetPosts(w http.ResponseWriter, r *http.Request) {

}

// Gets the specified post by its id
func GetPostByID(w http.ResponseWriter, r *http.Request) {}

// Updates the specified post by its id
func UpdatePost(w http.ResponseWriter, r *http.Request) {}

// Deletes the specified post by its id
func DeletePost(w http.ResponseWriter, r *http.Request) {}
