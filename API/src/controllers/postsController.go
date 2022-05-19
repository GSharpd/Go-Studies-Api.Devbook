package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	if err = post.Prepare(); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

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
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPostsRepository(db)

	posts, err := repo.GetPostsForUser(tokenUserID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, posts)
}

// Gets the specified post by its id
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repo := repositories.NewPostsRepository(db)

	post, err := repo.GetPostByID(postID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, post)
}

// Updates the specified post by its id
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repo := repositories.NewPostsRepository(db)

	existinPost, err := repo.GetPostByID(postID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if existinPost.PosterID != tokenUserID {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("you cannot edit someone else's post"))
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

	if err := post.Prepare(); err != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.UpdatePost(postID, post); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusOK, post)
}

// Deletes the specified post by its id
func DeletePost(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repo := repositories.NewPostsRepository(db)

	existinPost, err := repo.GetPostByID(postID)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if existinPost.PosterID != tokenUserID {
		responses.ErrorResponse(w, http.StatusForbidden, errors.New("you cannot edit someone else's post"))
		return
	}

	if err := repo.DeletePost(postID); err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSONResponse(w, http.StatusNoContent, nil)
}
