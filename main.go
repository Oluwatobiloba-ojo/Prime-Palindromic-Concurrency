package main

import (
	"awesomeProject/functions"
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println(err)
		return
	}

	if number < 1 || number > 50 {
		fmt.Println("Number must be between 1 and 50")
		return
	}

	sum := functions.FindAndSumUpPrimePalindromic(number)

	fmt.Println("The sum is ", sum)
}
