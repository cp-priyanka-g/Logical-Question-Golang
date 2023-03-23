package main

import "fmt"

//Question 6:Develop a program that takes in a string as input and returns True if the string is a palindrome and False otherwise

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

//Question 7: Develop a program that takes in a list of integers as input and returns a new list containing only the even integers from the input list
func EvenIntegersSlice() []int {
	numbers := [8]int{12, 111, 225, 4, 6, 88, 1, 9}
	var evenNumbers []int

	for i := 0; i < len(numbers); i++ {

		if numbers[i]%2 == 0 {

			evenNumbers = append(evenNumbers, numbers[i])

		}
	}
	return evenNumbers
}

func main() {
	//var Name string
	// fmt.Println("Enter String")
	// fmt.Scanln(&Name)
	// fmt.Println("Is it palindrome string", isPalindrome(Name))
	fmt.Println("Even Numbers:", EvenIntegersSlice())

}
