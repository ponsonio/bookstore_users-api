package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jcabrera/bookstore_users-api/domain/users"
	"github.com/jcabrera/bookstore_users-api/services"
	"github.com/jcabrera/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restError := errors.NewBadRequestError("bad user id")
		c.JSON(restError.Status, restError)
		return
	}

	user, saveErr := services.GetUser(userId)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context){
	var user users.User

	if restErr := c.ShouldBindJSON(&user); restErr != nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "not yet dude")
}