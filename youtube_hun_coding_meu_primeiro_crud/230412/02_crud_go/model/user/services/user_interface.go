package services

import (
	"crud_go/config/rest_err"
	model "crud_go/model/user"
)

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

type userDomainService struct {
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
