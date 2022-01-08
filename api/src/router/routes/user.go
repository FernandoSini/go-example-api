package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUser = []Route{

	{
		URI:                "/user",
		Metodo:             http.MethodPost,
		Function:           controllers.CreateUser,
		NeedAuthentication: false,
	}, {
		URI:                "/users",
		Metodo:             http.MethodGet,
		Function:           controllers.FindingUsers,
		NeedAuthentication: true,
	}, {
		URI:                "/user/{userId}",
		Metodo:             http.MethodGet,
		Function:           controllers.FindUser,
		NeedAuthentication: true,
	}, {
		URI:                "/user/{userId}",
		Metodo:             http.MethodPut,
		Function:           controllers.UpdateUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/user/{userId}",
		Metodo:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/user/{userId}/follow",
		Metodo:             http.MethodPost,
		Function:           controllers.FollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/user/{userId}/unfollow",
		Metodo:             http.MethodPost,
		Function:           controllers.UnfollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/user/{userId}/followers",
		Metodo:             http.MethodGet,
		Function:           controllers.FindFollowers,
		NeedAuthentication: true,
	}, {
		URI:                "/user/{userId}/following",
		Metodo:             http.MethodGet,
		Function:           controllers.UsersFollowing,
		NeedAuthentication: true,
	},
	{
		URI:                "/user/{userId}/password/update",
		Metodo:             http.MethodPost,
		Function:           controllers.UpdatePassword,
		NeedAuthentication: true,
	},
}
