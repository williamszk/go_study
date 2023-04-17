package user

import (
	"crud_go/config/rest_err"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Monkey! this is just test")
	c.JSON(err.Code, err)
}
