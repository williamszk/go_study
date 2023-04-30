package model

// The Controller will interact with the domain using interfaces or functions that return
// interfaces. The controller should not have direct access to the Model's structs.
// (I'm not sure why this would make total sense)
type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}


