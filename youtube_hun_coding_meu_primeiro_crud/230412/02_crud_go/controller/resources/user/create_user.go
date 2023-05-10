package user

import (
	"crud_go/config/logger"
	"crud_go/config/validation"
	"crud_go/controller/model/request"
	model_user "crud_go/model/user"
	"crud_go/view"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	tagJourney          = zap.String("journey", "createUser")
	UserDomainInterface model_user.UserDomainInterface
)

func (u userController) CreateUser(c *gin.Context) {
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

	// we create a new service here
	// service := service.NewUserDomainService()

	if err := u.serviceInterface.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully.", tagJourney)
	response := view.ConvertDomainToResponse(domain)
	logger.Info(fmt.Sprintf("%v", response), tagJourney)

	c.JSON(http.StatusOK, response)
}
