package user

import (
	"crud_go/config/validation"
	"crud_go/controller/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// err := rest_err.NewBadRequestError("Monkey! this is just test")
	// c.JSON(err.Code, err)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		// errRest := rest_err.NewBadRequestError(fmt.Sprintf("Sorry, something is wrong... %s", err.Error()))

		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)

}
