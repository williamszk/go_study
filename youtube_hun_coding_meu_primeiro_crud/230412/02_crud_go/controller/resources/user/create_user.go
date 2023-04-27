package user

import (
	"crud_go/config/logger"
	"crud_go/config/validation"
	"crud_go/controller/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	tagJourney = zap.String("journey", "createUser")
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", tagJourney)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info.", err,
			zap.Field{
				Key:    "journey",
				String: "createUser",
			})
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// response := response.UserResponse{
	// 	Id:    "test",
	// 	Email: userRequest.Email,
	// 	Name:  userRequest.Name,
	// 	Age:   userRequest.Age,
	// }

	logger.Info("User created successfully.", tagJourney)

	// c.JSON(http.StatusOK, response)
}
