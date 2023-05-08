package model

import (
	"crypto/md5"
	"encoding/hex"
)

// The Controller will interact with the domain using interfaces or functions that return
// interfaces. The controller should not have direct access to the Model's structs.
// (I'm not sure why this would make total sense)
type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	EncryptPassword()
}

func NewUserDomain(
	email,
	password,
	name string,
	age int8,
) UserDomainInterface {

	return &userDomain{
		email, password, name, age,
	}
}

func (u *userDomain) GetEmail() string {
	return u.email
}
func (u *userDomain) GetPassword() string {
	return u.password
}
func (u *userDomain) GetName() string {
	return u.name
}
func (u *userDomain) GetAge() int8 {
	return u.age
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}
