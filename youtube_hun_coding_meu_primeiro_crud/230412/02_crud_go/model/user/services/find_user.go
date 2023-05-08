package services

import (
	"crud_go/config/rest_err"
	model "crud_go/model/user"
)

func (u *userDomainService) FindUser(
	string) (*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
