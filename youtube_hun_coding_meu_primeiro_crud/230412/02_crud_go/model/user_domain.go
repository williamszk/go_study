package model

import (
	"crud_go/config/rest_err"
	"crypto/md5"
	"encoding/hex"
)

// The Controller will interact with the domain using interfaces or functions that return
// interfaces. The controller should not have direct access to the Model's structs.
// (I'm not sure why this would make total sense)
type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func NewUserDomain(
	email,
	password,
	name string,
	age int8,
) UserDomainInterface {

	return &UserDomain{
		email, password, name, age,
	}
}

func (u *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
