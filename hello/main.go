package main

import (
	"errors"
	"fmt"
)

func main() {
	var username = "nobody"

	fmt.Print("plese type your name")

	fmt.Scanln(&username)
	err := validate(username)
	if err != nil {
		fmt.Printf("You are not as cool as leslie! %s\n", username)
		return
	}
	fmt.Println("hello, ", username)
}
func validate(username string) error {
	if username != "leslie" {
		return errors.New("you are not leslie")
	}
	return nil

}
