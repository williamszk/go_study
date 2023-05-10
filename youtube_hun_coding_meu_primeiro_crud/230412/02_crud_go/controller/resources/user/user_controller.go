package user

import (
	"crud_go/model/user/service"

	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userController struct {
	serviceInterface service.UserDomainService
}

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userController{
		serviceInterface: serviceInterface,
	}
}
