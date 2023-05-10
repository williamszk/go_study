package service

import (
	"crud_go/config/rest_err"
	model "crud_go/model/user"
)

// The service is responsible for managing resources. Managing in the sense
// that it will be responsible for creating, updating, finding and deleting
// some resource.
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
