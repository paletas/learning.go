package main

import "fmt"

func main() {

	var exercisesAvailable = []string{"calculator", "helloworld"}

	fmt.Println("Exercises available:")

	for i, exercise := range exercisesAvailable {
		fmt.Printf("%d. %s\r\n", i+1, exercise)
		i++
	}

	var exerciseNumber int
	fmt.Println("Which exercise would you like to run?")
	fmt.Scanln(&exerciseNumber)

	if (exerciseNumber < 1) || (exerciseNumber > len(exercisesAvailable)) {
		fmt.Println("Invalid exercise number")
		return
	}

	switch exerciseNumber {
	case 1:
		useCalculator()
	case 2:
		useHelloWorld()
	}
}
