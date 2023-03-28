package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	if size == 0 {
		panic("Size cannot be zero")
	}

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

// Question 4: Develop a program that takes in a list of integers as input and returns the largest integer in the list.
func LargestIntegerNumber() float64 {
	var num [50]float64
	var total int
	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&total)
	if total == 0 {
		panic("Size cannot be zero")
	}
	for i := 0; i < total; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scan(&num[i])
	}
	for j := 1; j < total; j++ {
		if num[0] < num[j] {
			num[0] = num[j]
		}

	}
	return num[0]
}

// Question 5 : Develop a program that takes in a list of strings as input and returns the number of unique strings in the list.

func UniqueStringOfSlice() int {
	size := 0
	count := 0
	fmt.Print("Number of elements n=")
	fmt.Scanln(&size)
	fmt.Println("Enter  elements")
	elements := make([]string, size)
	for i := 0; i < size; i++ {

		fmt.Scanln(&elements[i])
	}
	if size == 0 {
		panic("Size cannot be zero")
	}

	PreviousElement := elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] != PreviousElement {
			count++
			PreviousElement = elements[i]
		}

	}
	return count
}

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
	var size int
	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&size)
	numbers := make([]int, size)
	var evenNumbers []int

	for i := 0; i < len(numbers); i++ {

		if numbers[i]%2 == 0 {

			evenNumbers = append(evenNumbers, numbers[i])

		}
	}
	return evenNumbers
}

//Question 8: Develop a program that takes in a list of integers as input and returns the second-largest integer in the list
func SecondLargestInteger() int {
	var size int

	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&size)
	intSlice := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scan(&intSlice[i])
	}

	fmt.Println("Before sort: ", intSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	return intSlice[1]
}

// Question 9: Develop a program that takes in a list of strings as input and returns a new list containing only the strings that start with a vowel
func VowelStringList() []string {
	//stringsList := []string{"Priya", "Manish", "Aryan", "Anvi", "eklavya", "urvashi"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter strings (one per line, press Enter to stop): ")

	stringsList := make([]string, 0)

	for {
		input, _ := reader.ReadString('\n')
		if input == "\n" {
			break
		}
		stringsList = append(stringsList, strings.TrimSpace(input))
	}
	vowelStringsList := []string{}

	for _, str := range stringsList {
		if strings.HasPrefix(str, "a") ||
			strings.HasPrefix(str, "e") ||
			strings.HasPrefix(str, "i") ||
			strings.HasPrefix(str, "o") ||
			strings.HasPrefix(str, "u") ||
			strings.HasPrefix(str, "A") ||
			strings.HasPrefix(str, "E") ||
			strings.HasPrefix(str, "I") ||
			strings.HasPrefix(str, "O") ||
			strings.HasPrefix(str, "U") {
			vowelStringsList = append(vowelStringsList, str)
		}
	}

	return vowelStringsList

}

//Question 10: Develop a program that takes in a list of integers as input and returns the sum of all integers in the list that are divisible by 3 or 5
func IsSumIntegerDivisible() int {
	var size int

	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&size)
	num := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scan(&num[i])
	}
	sum := 0

	for _, num := range num {
		if num%3 == 0 || num%5 == 0 {
			sum += num
		}
	}

	return sum
}

// Question 11: program to reverse string
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

	// Question 3
	fmt.Print("Shortest String in slice  : ", ShortestStringOfSlice())
	//Question 4
	fmt.Print("The largest number is : ", LargestIntegerNumber())
	//Question 5
	fmt.Print("The Number of Unique String is : ", UniqueStringOfSlice())
	//Question 6
	var Name string
	fmt.Println("Enter String")
	fmt.Scanln(&Name)
	fmt.Println("Is it palindrome string", isPalindrome(Name))
	//Question 7
	fmt.Println("Even Numbers:", EvenIntegersSlice())
	//Question 8
	fmt.Println("Second largest Integer", SecondLargestInteger())
	//Question 9
	fmt.Println("String with Vowels", VowelStringList())
	//Question 10
	fmt.Println("Sum of integer Divisible by 5 and 3:", IsSumIntegerDivisible())
	//Question 11
	fmt.Println("Reverse String:", ReverseString())

}
