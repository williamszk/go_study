package routes

import (
	resources_user "crud_go/controller/resources/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId",
		resources_user.FindUserById)
	r.GET("/getUserByEmail/:userEmail",
		resources_user.FindUserByEmail)
	r.POST("/createUser",
		resources_user.CreateUser)
	r.PUT("/updateUser/:userId",
		resources_user.UpdateUser)
	r.DELETE("/deleteUser/:userId",
		resources_user.DeleteUser)
}

