package routes

import (
	resource_user "crud_go/controller/resources/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController resource_user.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
