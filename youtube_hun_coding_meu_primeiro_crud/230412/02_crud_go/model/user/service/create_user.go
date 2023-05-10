package service

import (
	"crud_go/config/logger"
	"crud_go/config/rest_err"
	model "crud_go/model/user"
	"fmt"

	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	fmt.Println("Before hash, userDomain.GetPassword():", userDomain.GetPassword())

	userDomain.EncryptPassword()

	fmt.Println("After hash, userDomain.GetPassword():", userDomain.GetPassword())

	return nil
}
