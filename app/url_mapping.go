package app

import (
	"github.com/jcabrera/bookstore_users-api/controllers/ping"
	"github.com/jcabrera/bookstore_users-api/controllers/user"
)

func mapUrls(){
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
}
