//Create a Program that is run in the terminal. Where the user is taught the whole intermediate section of the go course. Thus add explanations and also examples of the code. This is a cool idea because when I get rusty on a section then I can just run the section that I would like to revise and then I will get the answers.

package main

import "fmt"

func main() {
	fmt.Println(".\n.\n.\nWelcome to My Golang Revision Tool.")
	fmt.Println("Last updated 05/09/2025")
	//Intro
	fmt.Println("Please select the chapter that you would like to revise")
	fmt.Print("1. All\n2. Closures\n3. Recursion\n100. Exit\nYour answer: ")
	var answer int
	fmt.Scan(&answer)

	//Switch over choices
	switch answer {
	case 1:
		fmt.Println("Revising all of the chapters")
	case 2:
		//closures
		fmt.Println("Revising Closures")
		counter := closureFunc()
		fmt.Println(counter())
		fmt.Println(counter())
		fmt.Println(counter())
		fmt.Println(counter())
	case 3:
		//recursion
		fmt.Println("Revising Recursion")
		fmt.Println(recursionFunc(10))
	case 4:
		fmt.Println("Revising ")
	default:
		fmt.Println("Exit")

	}

}

// A closure in Go is used when you need to capture and use variables from an outer function within an inner function, often for things like callbacks, maintaining state, or managing resource cleanup.
func closureFunc() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Recursion function of n + (n-1)
func recursionFunc(n int) int {
	if n < 1 {
		return n
	}
	return n + recursionFunc(n-1)
}
