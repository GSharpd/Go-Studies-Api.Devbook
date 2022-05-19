package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.GetPosts,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPostByID,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}/posts",
		Method:                 http.MethodGet,
		Function:               controllers.GetPostsByUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},
}
