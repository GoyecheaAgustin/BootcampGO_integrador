package main

import (
	"avatar/pkg/avatar"
	"fmt"
)

func main() {
	var username string
	println("Enter your username: ")
	fmt.Scanln(&username)
	avatar.AvatarGenerator().GenerateAndSaveAvatar(avatar.Information{UserName: username})
}
