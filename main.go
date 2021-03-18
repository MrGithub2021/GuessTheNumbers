package main

import (
	"fmt"
)

func add(a, b, c int) int {
	return a + b + c
}
func sub(a, b, c int) int {
	return a - b - c
}
func mult(a, b, c int) int {
	return a * b * c
}
func div(a, b, c int) int {
	return a / b / c
}
func re(a, b, c int) int {
	return a % b % c
}
func main() {
	fmt.Println("The three numbers, when added up, are:")
	fmt.Println(add(10, 2, 5))
	fmt.Println("The three numbers, when subtracted, are:")
	fmt.Println(sub(10, 2, 5))
	fmt.Println("The three numbers, when multiplied together, are:")
	fmt.Println(mult(10, 2, 5))
	fmt.Println("The three numbers, when divided, are:")
	fmt.Println(div(10, 2, 5))
	fmt.Println("What is the first number?")
	var first string
	fmt.Scanln(first)
	if first == "10" {
		fmt.Println("Yay! You got it right!")
		var second string
		fmt.Println("What is the second number?")
		fmt.Scanln(second)
		if second == "5" {
			fmt.Println("Yay! You got it!")
			var third string
			fmt.Scanln(third)
			if third == "2" {
				fmt.Println("Yay!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			} else {
				fmt.Println("Sorry!")
			}
		} else {
			fmt.Println("Try Again!")
		}
	} else {
		fmt.Print("Sorry!")
	}
}
