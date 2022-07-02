package interfaces

import "fmt"

type User struct {
	ID   string
	Name string
}

type UserMethods interface {
	GetUserById(id string)
	InsertUser(name string)
}

func (u *User) GetUserById(id string)  {}
func (u *User) InsertUser(name string) {}

type UserMock struct{}

func (u *UserMock) GetUserById(id string)  {}
func (u *UserMock) InsertUser(name string) {}

func test() {
	u := &User{}
	// why define an object of type User using & or without?
	// u1 := User{}

	var userMethods UserMethods

	userMethods = u

	fmt.Println(userMethods)
}
