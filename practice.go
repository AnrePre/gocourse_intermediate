//Create a Program that is run in the terminal. Where the user is taught the whole intermediate section of the go course. Thus add explanations and also examples of the code. This is a cool idea because when I get rusty on a section then I can just run the section that I would like to revise and then I will get the answers.

package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	//Intro
	fmt.Println("Welcome to Intermediate Go revision")
	fmt.Println("Please select the chapter that you would like to revise")
	fmt.Print("1. All\n2. Closures\n3. Recursion\n100. Exit\n")

	//closures
	counter := closureFunc()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	//recursion
	fmt.Println(recursionFunc(10))

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
