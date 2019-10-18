package entity

import(
	// "fmt"
)

type User struct{
	Username string
	Password string
	Email string
	Telephone string
}

func CreateUser(username, password, email, telephone string) User{
	user := User {username, password, email, telephone}
	return user
}

func (user User) GetUsername() string {  
    return user.Username
}

func (user User) GetPassword() string {  
    return user.Password
}

func (user User) GetEmail() string {  
    return user.Email
}

func (user User) GetTelephone() string {  
    return user.Telephone
}