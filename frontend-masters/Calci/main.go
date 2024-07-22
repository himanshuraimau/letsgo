package main

import "fmt"

func main() {
	defer fmt.Println("This is the end")

	var operations string
	var num1, num2 int

	fmt.Println("Calculator in Go")
	fmt.Println("====================")
	fmt.Println("Enter the operation you want to perform (Add, Subtract, Multiply, Divide): ")
	fmt.Scanf("%s", &operations)
    fmt.Println("Enter the first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanf("%d", &num2)
	switch operations {
	     case "Add":
			fmt.Println("The sum of the two numbers is: ", num1+num2)
		case "Subtract":
			fmt.Println("The difference of the two numbers is: ", num1-num2)
		case "Multiply":
			fmt.Println("The product of the two numbers is: ", num1*num2)
		case "Divide":
			fmt.Println("The division of the two numbers is: ", num1/num2)
		default:
			fmt.Println("Invalid operation")
	}

}
