package main

import "fmt"

func isPalindrome(str string) (isValid bool) {
	var result string

	for _, v := range str {
		result = string(v) + result
		fmt.Println("value string", result)
	}
	fmt.Println(str)
	fmt.Println(result)

	return str == result
}

func main() {
	var Name string
	fmt.Println("Enter String")
	fmt.Scanln(&Name)
	fmt.Println("Is it palindrome string", isPalindrome(Name))
}
