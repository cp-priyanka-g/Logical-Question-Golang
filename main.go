package main

import (
	"fmt"
)

func ReverseString() string {
	var Name string
	fmt.Println("Enter String")
	fmt.Scanln(&Name)

	var result string

	for _, v := range Name {
		result = string(v) + result
	}
	fmt.Println("Your String", Name)
	return result
}

func main() {
	fmt.Println("Reverse String:", ReverseString())
}
