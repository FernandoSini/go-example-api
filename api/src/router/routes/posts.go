package routes

import (
	"api/src/controllers"
	"net/http"
)

var routePosts = []Route{
	{
		URI:                "/post",
		Metodo:             http.MethodPost,
		Function:           controllers.CreatePosts,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts",
		Metodo:             http.MethodGet,
		Function:           controllers.FindPosts,
		NeedAuthentication: true,
	},
	{
		URI:                "/post/{postId}",
		Metodo:             http.MethodGet,
		Function:           controllers.FindPost,
		NeedAuthentication: true,
	},
	{
		URI:                "/post/{postId}",
		Metodo:             http.MethodPut,
		Function:           controllers.UpdatePost,
		NeedAuthentication: true,
	}, {
		URI:                "/post/{postId}",
		Metodo:             http.MethodDelete,
		Function:           controllers.DeletePost,
		NeedAuthentication: true,
	}, {
		URI:                "/user/{userId}/posts",
		Metodo:             http.MethodGet,
		Function:           controllers.FindPostsByUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/post/{postId}/like",
		Metodo:             http.MethodPost,
		Function:           controllers.LikePost,
		NeedAuthentication: true,
	}, {
		URI:                "/post/{postId}/unlike",
		Metodo:             http.MethodPost,
		Function:           controllers.UnlikePost,
		NeedAuthentication: true,
	},
}
