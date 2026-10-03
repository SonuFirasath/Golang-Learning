package main

import (
	"fmt"
	"github.com/SonuFirasath/Golang-Learning.git/auth"
	"github.com/SonuFirasath/Golang-Learning.git/user"
)

func main() {
	auth.LoginWihCrediantials("Firasath","Firasath@2004")

	session := auth.GetSession()

	fmt.Println(session)

	user := user.User{
		Email: "user@gmail.com",
		Name: "New User",
	}

	fmt.Println(user.Email, user.Name)

}