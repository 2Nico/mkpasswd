package main

import (
	"fmt"
	s "mkpasswd/src"
)

func main() {
	var counter uint64
	password := new(s.Password)

	s.GetParameter(password)
	s.HandleInputError(password.Length, password.Number, password.Character)

	for counter = 1; counter <= password.Quantity; counter++ {
		fmt.Println(s.RandomString(password.Length, password.Number, password.Character))
	}
}
