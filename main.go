package main

import (
	"fmt"
	"strconv"
)

// Question 1: Develop a program that takes in a list of integers as input and returns the sum of all even integers in the list(handle invalid input,
// non-integer values in the list) in golang

func TotalEvenNumbers(numbers []string) int {
	sum := 0
	for _, num := range numbers {
		input, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("%s is not an integer\n", num)
			continue
		}
		if input%2 == 0 {
			sum += input
		}
	}

	return sum
}

//Develop a program that takes in a string as input and returns the number of vowels in the string.

func CountVowels(name string) int {
	count := 0
	for _, i := range name {
		switch i {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}

	return count
}

func main() {
	var name string
	fmt.Println("Question 1")
	// Question 1
	numbers := []string{"1", "bananana", "3", "four", "5", "22", "2"}
	Sum := TotalEvenNumbers(numbers)
	fmt.Println(Sum)

	// Question 2
	fmt.Println("Question 2")
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("The number of vowels in '%s' is %d\n", name, CountVowels(name))

}
