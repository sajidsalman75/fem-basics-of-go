package main

import (
	"fmt"
)

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Calculator GO 1.0")
	fmt.Println("=================")
	fmt.Println("Which operation do you want to perform? (add, sub, mul, div)")
	fmt.Scanf("%s\n", &operation)
	fmt.Println("Enter first number")
	fmt.Scanf("%d\n", &number1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d\n", &number2)

	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "sub":
		fmt.Println(number1 - number2)
	case "mul":
		fmt.Println(number1 * number2)
	case "div":
		fmt.Println(number1 / number2)
	}
}
