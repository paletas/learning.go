package main

import "fmt"

func useCalculator() {

	fmt.Println("What operation would you like to perform?")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	var operation int
	fmt.Scanln(&operation)

	fmt.Println("What is the first number?")
	var firstNumber float64
	fmt.Scanln(&firstNumber)

	fmt.Println("What is the second number?")
	var secondNumber float64
	fmt.Scanln(&secondNumber)

	var result float64

	switch operation {
	case 1:
		fmt.Println("You chose addition")
		result = firstNumber + secondNumber
	case 2:
		fmt.Println("You chose subtraction")
		result = firstNumber - secondNumber
	case 3:
		fmt.Println("You chose multiplication")
		result = firstNumber * secondNumber
	case 4:
		fmt.Println("You chose division")
		result = firstNumber / secondNumber
	default:
		fmt.Println("Invalid operation")
	}

	fmt.Println("The result is", result)
}
