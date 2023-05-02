package model

import (
	"crud_go/config/logger"
	"crud_go/config/rest_err"

	"go.uber.org/zap"
)

func (u *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init create user model", zap.String("journey", "createUser"))

	u.EncryptPassword()
	return nil
}
