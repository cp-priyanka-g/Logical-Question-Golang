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

func main() {
	fmt.Println("Hello World!")
	// Question 1
	numbers := []string{"1", "bananana", "3", "four", "5", "22", "2"}

	Sum := TotalEvenNumbers(numbers)
	fmt.Println(Sum)

}
