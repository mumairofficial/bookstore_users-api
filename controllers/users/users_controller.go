package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mumairofficial/bookstore_users-api/domain/users"
	"github.com/mumairofficial/bookstore_users-api/services"
	"github.com/mumairofficial/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {}

func SearchUser(c *gin.Context) {}
