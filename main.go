package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

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
func main() {
	//var Name string
	// fmt.Println("Enter String")
	// fmt.Scanln(&Name)
	// fmt.Println("Is it palindrome string", isPalindrome(Name))
	//fmt.Println("Even Numbers:", EvenIntegersSlice())
	//fmt.Println("Second largest Integer", SecondLargestInteger())
	//fmt.Println("String with Vowels", VowelStringList())
	fmt.Println("String with Vowels", IsSumIntegerDivisible())

}
