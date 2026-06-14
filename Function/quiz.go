package main

import (
	"errors"
	"fmt"
)

func Quiz() {
	scores := []int{10, 5, 8, 9, 7}

	total := sum(scores)
	fmt.Println(total)

	calculator, err := calculator(5, 5, "*")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(calculator)

}

func sum(scores []int) int {
	var result int
	for _, sum := range scores {
		result = result + sum
	}

	return result
}

func calculator(num1, num2 int, operator string) (int, error) {

	var result int
	var errResult error

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		errResult = errors.New("Unknown operator")
	}

	return result, errResult
}
