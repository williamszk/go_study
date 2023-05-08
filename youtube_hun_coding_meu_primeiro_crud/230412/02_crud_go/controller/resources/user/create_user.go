package user

import (
	"crud_go/config/logger"
	"crud_go/config/validation"
	"crud_go/controller/model/request"
	model_user "crud_go/model/user"
	"crud_go/model/user/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	tagJourney          = zap.String("journey", "createUser")
	UserDomainInterface model_user.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", tagJourney)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validate user info.",
			err,
			tagJourney)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model_user.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	// if err := domain.CreateUser(); err != nil {
	// 	c.JSON(err.Code, err)
	// 	return
	// }

	service := services.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	// response := response.UserResponse{
	// 	Id:    "test",
	// 	Email: userRequest.Email,
	// 	Name:  userRequest.Name,
	// 	Age:   userRequest.Age,
	// }

	logger.Info("User created successfully.", tagJourney)

	c.String(http.StatusOK, "")
}
