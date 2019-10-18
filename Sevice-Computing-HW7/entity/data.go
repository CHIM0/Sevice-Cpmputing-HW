package entity

import (
	"encoding/json"
	"fmt"
	"os"
)
var users []User
var curUser []User

var usersPath = "database/Users.txt"
var curUserPath = "database/curUser.txt"

func init(){
	LoadUsers("users")
	LoadUsers("curUser")
}

func isUserExited(username string) bool{
	for _, item := range users {
		if item.GetUsername() == username {
			return true
		}
	}
	return false
}

func checkPassword(username, password string) bool{
	for _, item := range users {
		if item.GetUsername() == username {
			if item.GetPassword() == password{
				return true
			}
			return false
		}
	}
	return false
}

func UserRegister(username, password, email, telephone string) bool{
	if (isUserExited(username)) {
		fmt.Println("ERROR: The Username is already exited!")
		return false
	}
	user := CreateUser(username, password, email, telephone)
	users = append(users, user)
	if SaveUsers("users") {
		return true
	}else {return false}
}

func UserSignIn(username, password string) bool {
	if len(curUser) > 0{
		if curUser[0].GetUsername() == username {
			fmt.Println("ERROR: You have already sign in.")
		}else{
			fmt.Println("ERROR:",curUser[0].GetUsername(),"has already sign in, please sign out first if you want to sign in.")
		}
		return false
	}
	if !isUserExited(username) {
		fmt.Println("ERROR: Wrong username or password!")
		return false
	}
	if !checkPassword(username, password) {
		fmt.Println("ERROR: Wrong username or password!")
		return false
	}
	user := CreateUser(username,"","","")
	curUser = append(curUser, user)
	if SaveUsers("curUser") {
		return true
	}else{return false}
}

func UserSignOut() bool{
	if len(curUser) == 0{
		fmt.Println("ERROR: No user signed in.")
		return false
	}else{
		curUser = curUser[0:0]
		if SaveUsers("curUser") {
			return true
		}else {return false}
	}
}

func LoadUsers(mode string) bool {
	path := ""
	switch mode{
	case "users":
		path = usersPath
		break
	case "curUser":
		path = curUserPath
	}
	if file, err := os.Open(path); err != nil{
		fmt.Println("Error: Can't reach database!")
		file.Close()
		return false
	}else{
		JSON := json.NewDecoder(file)
		switch mode{
		case "users":
			if err := JSON.Decode(&users); err !=nil {
				fmt.Println("Error: Database injured!")
				return false
			}
			break
		case "curUser":
			if err := JSON.Decode(&curUser); err !=nil {
				fmt.Println("Error: Database injured!")
				return false
			}
		}
		return true
	}
}

func SaveUsers(mode string) bool{
	path := ""
	switch mode{
	case "users":
		path = usersPath
		break
	case "curUser":
		path = curUserPath
	}
	if file, err := os.Create(path); err != nil {
		fmt.Println("Error: Can't reach database!")
		file.Close()
		return false
	}else{
		JSON := json.NewEncoder(file)
		switch mode{
		case "users":
			if err := JSON.Encode(&users); err !=nil {
				fmt.Println("Error: Database injured!")
				return false
			}
			break
		case "curUser":
			if err := JSON.Encode(&curUser); err !=nil {
				fmt.Println("Error: Database injured!")
				return false
			}
		}
		return true
	}
}