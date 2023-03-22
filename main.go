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

//Question 2: Develop a program that takes in a string as input and returns the number of vowels in the string.

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

// Question 4: Develop a program that takes in a list of integers as input and returns the largest integer in the list.
func LargestIntegerNumber() float64 {
	var num [50]float64
	var total int
	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&total)
	for i := 0; i < total; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scan(&num[i])
	}
	for j := 1; j < total; j++ {
		if num[0] > num[j] {
			num[0] = num[j]
		}

	}
	return num[0]
}

//Question 3: Develop a program that takes in a list of strings as input and returns the shortest string in the list.
func ShortestStringOfSlice() string {
	size := 0
	fmt.Print("Number of elements n=")
	fmt.Scanln(&size)
	fmt.Println("Enter  elements")
	elements := make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&elements[i])
	}
	fmt.Println("Entered Array of elements:", elements)

	smallestElement := len(elements[0])
	index := 0
	for i := 1; i < len(elements); i++ {
		if len(elements[i]) < smallestElement {
			index = i
			smallestElement = len(elements[i])
		}

	}
	return elements[index]
}

func main() {
	// var name string
	// fmt.Println("Question 1")
	// // Question 1
	// numbers := []string{"1", "bananana", "3", "four", "5", "22", "2"}
	// Sum := TotalEvenNumbers(numbers)
	// fmt.Println(Sum)

	// // Question 2
	// fmt.Println("Question 2")
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)
	//fmt.Printf("The number of vowels in '%s' is %d\n", name, CountVowels(name))

	// Question 3
	fmt.Print("Shortest String in slice  : ", ShortestStringOfSlice())
	//Question 4

	//fmt.Print("The largest number is : ", LargestIntegerNumber())
}
